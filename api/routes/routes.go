package routes

import (
	"api_pattern_go/api/controller/teste"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/teste", teste.Criar)
	router.GET("/teste/:id", teste.Visualizar)
}
