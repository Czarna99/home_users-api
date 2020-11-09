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

func GetGroup(c *gin.Context) {
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

func CreateGroup(c *gin.Context) {
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
