package routes

import (
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
		usuarioGroup := main.Group("/usuarios")
		{
			usuarioRoutes(usuarioGroup)
		}
		testeGroup := main.Group("/testes", middleware.Auth())
		{
			testeRoutes(testeGroup)
		}
		loginGroup := main.Group("/login")
		{
			loginRoutes(loginGroup)
		}
		permissaoGroup := main.Group("/permissoes", middleware.Auth())
		{
			permissaoRoutes(permissaoGroup)
		}
	}

	return router
}
