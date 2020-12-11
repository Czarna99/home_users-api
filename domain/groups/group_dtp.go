package groups

import (
	"strings"

	"github.com/Pawelek242/home_users-api/utils/errors"
)

//Group - groups for user
type Group struct {
	ID          int64  `json:"id"`
	GroupName   string `json:"group_name"`
	Privileges  string `json:"privileges"`
	DateCreated string `json:"date_created"`
}

//Validate - group validation
func (groups *Group) Validate() *errors.RestErr {
	var error []string
	groups.GroupName = strings.TrimSpace(strings.ToLower(groups.GroupName))

	if groups.GroupName == "" {
		return errors.NewBadRequest(append(error, "Invalid group name"))
	}
	return nil
}
