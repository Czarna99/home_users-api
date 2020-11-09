package users

import (
	"fmt"
	"strings"

	"github.com/Pawelek242/home_users-api/dataresources/mysql/users_db"
	"github.com/Pawelek242/home_users-api/utils/date_utils"
	"github.com/Pawelek242/home_users-api/utils/errors"
)

const (
	errorNoRows      = "no rows in result set"
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

var (
//usersDB = make(map[int64]*User)
)

//Get - get user from database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFound(
				fmt.Sprintf("user %d not found", user.ID))

		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()))

	}
	return nil
}

//Save - Saving user data into database
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequest(
				fmt.Sprintf("email %s is already used", user.Email)) //TODO - mail verify (correct email syntax)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s ", err.Error()))
	}

	user.ID = userID

	return nil
}

//Update - update user
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err)) //TODO mysql_utils - error MYSQL error handling
	}
	return nil

}
