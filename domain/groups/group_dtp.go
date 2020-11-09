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

func (groups *Group) Validate() *errors.RestErr {
	groups.GroupName = strings.TrimSpace(strings.ToLower(groups.GroupName))

	if groups.GroupName == "" {
		return errors.NewBadRequest("Invalid group name")
	}
	return nil
}
