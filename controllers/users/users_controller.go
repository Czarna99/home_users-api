package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pawelek242/home_users-api/domain/users"
	"github.com/Pawelek242/home_users-api/services"
	"github.com/Pawelek242/home_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("invalid user id - should be a number")
		c.JSON(err.Code, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid data")
		c.JSON(restErr.Code, restErr)
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

//UpdateUser - function for update user
func UpdateUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("invalid user id - should be a number")
		c.JSON(err.Code, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid data")
		c.JSON(restErr.Code, restErr)
		return
	}
	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
