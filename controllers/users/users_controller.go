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

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	var error []string
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequest(append(error, "invalid user id - should be a number"))

	}
	return userID, nil
}
func Get(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Code, idErr)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func Create(c *gin.Context) {

	var (
		user  users.User
		error []string
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest(append(error, "Invalid data"))
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
func Update(c *gin.Context) {
	var error []string
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Code, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest(append(error, "Invalid data"))
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

func Delete(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Code, idErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, users)

}
