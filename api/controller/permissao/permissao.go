package permissao

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/enum"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/repository"
	"api_pattern_go/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var p models.Permissao

	if err = ginctx.ShouldBindJSON(&p); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewPermissaoRepository(dbConetion.DB).Create(&p); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, p))
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
