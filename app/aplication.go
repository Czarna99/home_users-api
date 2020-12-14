package app

import (
	"github.com/Pawelek242/home_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication lol
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8080")
}
