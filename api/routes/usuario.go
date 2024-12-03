package routes

import (
	"api_pattern_go/api/controller/usuario"
	"api_pattern_go/api/middleware"
	"github.com/gin-gonic/gin"
)

func usuarioRoutes(r *gin.RouterGroup) {
	r.POST(route, usuario.Criar)
	r.GET(routeId, middleware.Auth(), usuario.Visualizar)
	r.GET(route, middleware.Auth(), usuario.Listar)
	r.GET(routeDropdown, middleware.Auth(), usuario.Dropdown)
	r.PUT(route, middleware.Auth(), usuario.Atualizar)
	r.DELETE(routeId, middleware.Auth(), usuario.Deletar)

	r.POST(routeId+"/permissao/:idPermissao", middleware.Auth(), usuario.AtribuirPermissao)
	r.DELETE(routeId+"/permissao/:idPermissao", middleware.Auth(), usuario.RemoverPermissao)
}
