package services

import (
	groups "github.com/Pawelek242/home_users-api/domain/groups"
	"github.com/Pawelek242/home_users-api/utils/errors"
)

func GetGroup(groupID int64) (*groups.Group, *errors.RestErr) {
	result := &groups.Group{ID: groupID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
func CreateGroup(groups groups.Group) (*groups.Group, *errors.RestErr) {
	if err := groups.Validate(); err != nil {
		return nil, err
	}

	if err := groups.Save(); err != nil {
		return nil, err
	}
	return &groups, nil
}
