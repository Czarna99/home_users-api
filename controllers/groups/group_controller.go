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
	groupID, groupErr := strconv.ParseInt(groupIDParam, 10, 64)
	if groupErr != nil {
		return 0, errors.NewBadRequest("invalid ID - Should be a number.")
	}
	return groupID, nil

}

func Get(c *gin.Context) {
	groupID, groupErr := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if groupErr != nil {
		err := errors.NewBadRequest("invalid group id - should be a number")
		c.JSON(err.Code, err)
		return
	}
	group, getErr := services.GetGroup(groupID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, group)

}

func Create(c *gin.Context) {
	var group groups.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		restErr := errors.NewBadRequest("Invalid data")
		c.JSON(restErr.Code, restErr)
	}
	result, saveErr := services.CreateGroup(group)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}

	fmt.Println(group)
	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {
	groupID, groupErr := strconv.ParseInt(c.Param("group_id"), 10, 64)
	fmt.Printf("%d", groupID)
	if groupErr != nil {
		err := errors.NewBadRequest("Invalid group ID")
		c.JSON(err.Code, err)
		fmt.Println("SSIJ")
		return
	}

	var group groups.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		restErr := errors.NewBadRequest("Invalid data")
		c.JSON(restErr.Code, restErr)

		return
	}
	group.ID = groupID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateGroup(isPartial, group)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
func Delete(c *gin.Context) {

	groupID, groupErr := getGroupID(c.Param("group_id"))
	if groupErr != nil {
		c.JSON(groupErr.Code, groupErr)
	}
	if err := services.DeleteGroup(groupID); err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
