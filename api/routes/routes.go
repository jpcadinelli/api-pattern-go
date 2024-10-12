package routes

import (
	"api_pattern_go/api/controller/teste"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/testes", teste.Criar)
	router.GET("/testes/:id", teste.Visualizar)
}
