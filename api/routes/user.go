package routes

import (
	"api_pattern_go/api/controller/user"
	"api_pattern_go/api/middleware"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.RouterGroup) {
	r.POST(route, user.Criar)
	r.GET(routeId, middleware.Auth(), user.Visualizar)
	r.POST(routeFiltro, middleware.Auth(), user.Listar)
	r.PUT(route, middleware.Auth(), user.Atualizar)
	r.DELETE(routeId, middleware.Auth(), user.Deletar)
}
