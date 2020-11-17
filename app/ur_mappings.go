package app

import (
	"github.com/Pawelek242/home_users-api/controllers/groups"
	"github.com/Pawelek242/home_users-api/controllers/ping"
	"github.com/Pawelek242/home_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	//router.GET("users/search", controllers.SearchUser)
	router.POST("/users", users.Create)
	router.POST("/groups", groups.Create)
	router.GET("/groups/:group_id", groups.Get)
	router.PATCH("/users/:user_id", users.Update)
	router.PATCH("/groups/:group_id", groups.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.DELETE("/groups/:group_id", groups.Delete)
}
