package routes

import (
	"api_pattern_go/api/controller/teste"
	"github.com/gin-gonic/gin"
)

func testeRoutes(r *gin.RouterGroup) {
	r.POST(route, teste.Criar)
	r.GET(routeId, teste.Visualizar)
	r.POST(routeFiltro, teste.Listar)
	r.PUT(route, teste.Atualizar)
	r.DELETE(routeId, teste.Deletar)
}
