package groups

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pawelek242/home_users-api/domain/groups"
	"github.com/Pawelek242/home_users-api/services"
	"github.com/Pawelek242/home_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getGroupID(groupIDParam string) (int64, *errors.RestErr) {
	var error []string
	groupID, groupErr := strconv.ParseInt(groupIDParam, 10, 64)
	if groupErr != nil {
		return 0, errors.NewBadRequest(append(error, "invalid ID - Should be a number."))
	}
	return groupID, nil

}

//GetGroup - function for get group from database
func GetGroup(c *gin.Context) {
	var error []string
	groupID, groupErr := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if groupErr != nil {
		err := errors.NewBadRequest(append(error, "invalid group id - should be a number"))
		c.JSON(err.Code, err)
		return
	}
	group, getErr := services.GroupService.GetGroup(groupID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, group)

}

// CreateGroup - function for create group in database
func CreateGroup(c *gin.Context) {
	var error []string
	var group groups.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		restErr := errors.NewBadRequest(append(error, "Invalid data"))
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.GroupService.CreateGroup(group)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}

	fmt.Println(group)
	c.JSON(http.StatusCreated, result)
}

//UpdateGroup - function for updating group information in database
func UpdateGroup(c *gin.Context) {
	var error []string
	groupID, groupErr := strconv.ParseInt(c.Param("group_id"), 10, 64)
	fmt.Printf("%d", groupID)
	if groupErr != nil {
		err := errors.NewBadRequest(append(error, "Invalid group ID"))
		c.JSON(err.Code, err)

		return
	}

	var group groups.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		restErr := errors.NewBadRequest(append(error, "Invalid data"))
		c.JSON(restErr.Code, restErr)

		return
	}
	group.ID = groupID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.GroupService.UpdateGroup(isPartial, group)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

//DeleteGroup - function for deleting group from database
func DeleteGroup(c *gin.Context) {

	groupID, groupErr := getGroupID(c.Param("group_id"))
	if groupErr != nil {
		c.JSON(groupErr.Code, groupErr)
	}
	if err := services.GroupService.DeleteGroup(groupID); err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
