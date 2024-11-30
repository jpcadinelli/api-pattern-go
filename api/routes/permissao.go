package routes

import (
	"api_pattern_go/api/controller/permissao"
	"github.com/gin-gonic/gin"
)

func permissaoRoutes(r *gin.RouterGroup) {
	r.POST(route, permissao.Criar)
	r.GET(routeId, permissao.Visualizar)
	r.POST(routeFiltro, permissao.Listar)
	r.PUT(route, permissao.Atualizar)
	r.DELETE(routeId, permissao.Deletar)
}
