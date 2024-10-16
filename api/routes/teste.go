package routes

import (
	"api_pattern_go/api/controller/teste"
	"github.com/gin-gonic/gin"
)

const (
	testes       = "/testes"
	testesId     = "/testes/:id"
	testesFiltro = "/testes/filtro"
)

func testeRoutes(r *gin.Engine) {
	r.POST(testes, teste.Criar)
	r.GET(testesId, teste.Visualizar)
	r.POST(testesFiltro, teste.Listar)
	r.PUT(testes, teste.Atualizar)
	r.DELETE(testesId, teste.Deletar)
}
