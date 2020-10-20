package users

import (
	"fmt"

	"github.com/Pawelek242/home_users-api/dataresources/mysql/users_db"
	"github.com/Pawelek242/home_users-api/utils/date_utils"
	"github.com/Pawelek242/home_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("User %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if usersDB[user.ID] != nil {
		if current.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("This email %s has been already registered", user.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("This user %d already exist", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.ID] = user
	return nil
}
