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
	queryGetGroup    = "SELECT id, group_name FROM group WHERE id=?"
	queryInsertGroup = "INSERT INTO users_db.groups(group_name, date_created) VALUES(?, ?);" //TODO users_db.groups change for groups
	indexUniqueName  = "group_name_UNIQUE"
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
	if err := result.Scan(&groups.ID, &groups.GroupName, &groups.DateCreated); err != nil {
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
	fmt.Println("jebany nie działa")
	groups.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(groups.GroupName, groups.DateCreated)
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
