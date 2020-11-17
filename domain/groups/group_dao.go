package groups

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
	queryGetGroup    = "SELECT id, group_name, privileges, date_created FROM users_db.groups WHERE id=?"
	queryInsertGroup = "INSERT INTO users_db.groups(group_name, date_created, privileges) VALUES(?, ?, ?);" //TODO users_db.groups change for groups
	indexUniqueName  = "group_name_UNIQUE"
	queryUpdateGroup = "UPDATE users_db.groups SET group_name=?, privileges=? WHERE id=?;"
	queryDeleteGroup = "DELETE FROM users_db.groups WHERE id=?"
)

var ()

//Get - get group from database
func (groups *Group) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetGroup)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(groups.ID)
	if err := result.Scan(&groups.ID, &groups.GroupName, &groups.Privileges, &groups.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFound(
				fmt.Sprintf("group %d not found", groups.ID))

		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", groups.ID, err.Error()))

	}
	return nil
}

//Save - Saving user data into database
func (groups *Group) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertGroup)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	groups.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(groups.GroupName, groups.DateCreated, groups.Privileges)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueName) {
			return errors.NewBadRequest(
				fmt.Sprintf("Group name %s already exist", groups.GroupName))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save group: %s", err.Error()))
	}

	groupID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save group: %s ", err.Error()))
	}

	groups.ID = groupID

	return nil
}

//Update - updating group info
func (groups *Group) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateGroup)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(groups.GroupName, groups.Privileges, groups.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err))
	}
	return nil
}
func (groups *Group) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteGroup)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(groups.ID); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err)) //TODO mysql_utils - error MYSQL error handling
	}
	return nil
}
