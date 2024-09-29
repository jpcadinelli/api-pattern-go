package routes

import (
	"api_pattern_go/api/controller/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/users", user.CreateUser)
	router.GET("/users", user.GetUsers)
}
