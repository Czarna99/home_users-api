package services

import (
	"fmt"

	groups "github.com/Pawelek242/home_users-api/domain/groups"
	"github.com/Pawelek242/home_users-api/utils/errors"
)

func GetGroup(groupID int64) (*groups.Group, *errors.RestErr) {
	result := &groups.Group{ID: groupID}
	if err := result.GetGroup(); err != nil {
		return nil, err
	}
	return result, nil
}
func CreateGroup(groups groups.Group) (*groups.Group, *errors.RestErr) {
	if err := groups.Validate(); err != nil {
		return nil, err
	}

	if err := groups.SaveGroup(); err != nil {
		return nil, err
	}
	return &groups, nil
}
func UpdateGroup(isPartial bool, groups groups.Group) (*groups.Group, *errors.RestErr) {
	current, err := GetGroup(groups.ID)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if groups.GroupName != "" {
			current.GroupName = groups.GroupName
		}
		if groups.Privileges != "" {
			current.Privileges = groups.Privileges
		}
	} else {
		if err := groups.Validate(); err != nil {
			return nil, err
		}
		current.GroupName = groups.GroupName
		current.Privileges = groups.Privileges
	}
	fmt.Printf("%s", current.GroupName)
	fmt.Printf("%s", current.Privileges)
	if err := current.UpdateGroup(); err != nil {
		return nil, err
	}
	return current, nil

}
func DeleteGroup(groupID int64) *errors.RestErr {
	group := &groups.Group{ID: groupID}
	return group.DeleteGroup()
}
