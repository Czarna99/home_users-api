package users

import (
	"fmt"
	"strings"

	"github.com/Pawelek242/home_users-api/dataresources/mysql/users_db"
	"github.com/Pawelek242/home_users-api/logger"
	"github.com/Pawelek242/home_users-api/utils/date_utils"
	"github.com/Pawelek242/home_users-api/utils/errors"
)

const (
	errorNoRows           = "no rows in result set"
	indexUniqueEmail      = "email_UNIQUE"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus = "SELECT id, first_name, last_Name, email, date_created, status FROM users WHERE status=?;"
)

var (
//usersDB = make(map[int64]*User)
)

//Get - get user from database
func (user *User) Get() *errors.RestErr {
	var error []string
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError(append(error, "Database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			logger.Error("error when trying to get user by ID", err)
			return errors.NewInternalServerError(append(error, "Database error"))
		}
	}
	return nil
}

//Save - Saving user data into database
func (user *User) Save() *errors.RestErr {
	var error []string

	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError(append(error, "Database error"))
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowDBFormat()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequest(append(error, fmt.Sprintf("email %s is already used", user.Email)))
		}
		logger.Error("error when trying save new user", err)
		return errors.NewInternalServerError(append(error, "Database error"))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError(append(error, "Database error"))
	}

	user.ID = userID

	return nil
}

//Update - update user
func (user *User) Update() *errors.RestErr {
	var error []string
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(append(error, err.Error()))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(append(error, fmt.Sprintf("%s", err))) //TODO mysql_utils - error MYSQL error handling
	}
	return nil

}

//Delete - deleting user
func (user *User) Delete() *errors.RestErr {
	var error []string
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(append(error, err.Error()))
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.ID); err != nil {
		return errors.NewInternalServerError(append(error, fmt.Sprintf("%s", err))) //TODO mysql_utils - error MYSQL error handling
	}
	return nil
}

//FindByStatus - looking for users with specific account status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	var error []string
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(append(error, err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(append(error, err.Error()))
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(append(error, err.Error()))
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFound(append(error, fmt.Sprintf("There is no users with status %s", status)))
	}
	return result, nil
}
