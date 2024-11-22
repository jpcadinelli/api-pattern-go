package routes

import (
	"api_pattern_go/api/controller/user"
	"github.com/gin-gonic/gin"
)

const (
	users       = "/"
	usersId     = "/:id"
	usersFiltro = "/filtro"
)

func userRoutes(r *gin.RouterGroup) {
	r.POST(users, user.Criar)
	r.GET(usersId, user.Visualizar)
	r.POST(usersFiltro, user.Listar)
	r.PUT(users, user.Atualizar)
	r.DELETE(usersId, user.Deletar)
}
