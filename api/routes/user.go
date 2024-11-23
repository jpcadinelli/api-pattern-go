package routes

import (
	"api_pattern_go/api/controller/user"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.RouterGroup) {
	r.POST(route, user.Criar)
	r.GET(routeId, user.Visualizar)
	r.POST(routeFiltro, user.Listar)
	r.PUT(route, user.Atualizar)
	r.DELETE(routeId, user.Deletar)
}
