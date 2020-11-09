package users

import (
	"strings"

	"github.com/Pawelek242/home_users-api/utils/errors"
)

//User - user struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Privileges  string `json:"privileges"`
}

//Global - global struct
type Global struct {
	updatedBy string
	createdBy string
	deletedBy string

	created string
	updated string
	deleted string
}

//Validate - validating data provided into database
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	if user.Email == "" {
		return errors.NewBadRequest("invalid email adress")
	}
	if user.FirstName == "" {
		return errors.NewBadRequest("Name field shouldn't be empty.")
	}
	if user.LastName == "" {
		return errors.NewBadRequest("Lastname field shouldn't be empty.")
	}
	return nil

}
