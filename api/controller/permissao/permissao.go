package permissao

import (
	"api_pattern_go/api/global"
	"api_pattern_go/api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	idUsuarioLogado, err := global.GetIdUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, idUsuarioLogado))
	return
}

func Visualizar(_ *gin.Context) {
	return
}

func Listar(_ *gin.Context) {
	return
}

func Atualizar(_ *gin.Context) {
	return
}

func Deletar(_ *gin.Context) {
	return
}
