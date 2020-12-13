package users

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Pawelek242/home_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

//User - user struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Privileges  string `json:"privileges"`
	Status      string `json:"status"`
	Password    string `json:"password"`
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
	var err []string
	if user.Email == "" {
		return errors.NewInternalServerError(append(err, "Email field cannot be empty"))
	}
	if user.FirstName == "" {
		return errors.NewInternalServerError(append(err, "Name field shouldn't be empty."))
	}
	if user.LastName == "" {
		return errors.NewInternalServerError(append(err, "Lastname field shouldn't be empty."))
	}
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	return nil
}
func (user *User) CheckPassword() *errors.RestErr {

	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var err []string

	user.Password = strings.TrimSpace(user.Password)
	for _, ch := range user.Password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
		case unicode.IsUpper(ch):
			uppercasePresent = true
		case unicode.IsLower(ch):
			lowercasePresent = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
		case ch == ' ':
		}
		passLen++
	}

	if !lowercasePresent {

		err = append(err, "Lowercase letter missing.")
	}
	if !uppercasePresent {
		err = append(err, "Uppercase letter missing.")
	}
	if !numberPresent {
		err = append(err, "At least one numeric character required.")
	}
	if !specialCharPresent {
		err = append(err, "Special character missing.") //TODO special characters jakie
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		err = append(err, fmt.Sprintf("Password length must be between %d to %d characters long.", minPassLength, maxPassLength))
	}
	if err != nil {
		return errors.PassError(err)
	}
	return nil

}
