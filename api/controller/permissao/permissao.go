package permissao

import (
	"api_pattern_go/api/controller/usuario"
	"api_pattern_go/api/global/enum"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := usuario.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !VerificaPermissaoAdminUsuario(*usuarioLogado) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, nil))
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

func VerificaPermissaoAdminUsuario(usuario models.UsuarioDTOResponse) bool {
	for _, permissao := range usuario.Permissoes {
		if permissao.Nome == enum.PermissaoSistemaAdmin {
			return true
		}
	}
	return false
}
