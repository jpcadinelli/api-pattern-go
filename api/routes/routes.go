package routes

import (
	"api_pattern_go/api/controller/login"
	"api_pattern_go/api/controller/teste"
	"api_pattern_go/api/controller/user"
	"api_pattern_go/api/middleware"
	"github.com/gin-gonic/gin"
)

const (
	route       = "/"
	routeId     = "/:id"
	routeFiltro = "/filtro"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		userGroup := main.Group("/users")
		{
			userGroup.POST(route, user.Criar)
			userGroup.GET(routeId, user.Visualizar)
			userGroup.POST(routeFiltro, user.Listar)
			userGroup.PUT(route, user.Atualizar)
			userGroup.DELETE(routeId, user.Deletar)
		}
		testeGroup := main.Group("/testes", middleware.Auth())
		{
			testeGroup.POST(route, teste.Criar)
			testeGroup.GET(routeId, teste.Visualizar)
			testeGroup.POST(routeFiltro, teste.Listar)
			testeGroup.PUT(route, teste.Atualizar)
			testeGroup.DELETE(routeId, teste.Deletar)
		}
		loginGroup := main.Group("/login")
		{
			loginGroup.POST(route, login.Login)
		}
	}

	return router
}
