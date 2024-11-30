package login

import (
	dbConection "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/repository"
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ginctx *gin.Context) {
	var l models.Login

	if err := ginctx.ShouldBindJSON(&l); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	usuario, err := repository.NewUsuarioRepository(dbConection.DB).FindByEmail(l.Email)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if usuario.Password != services.SHA256Encoder(l.Password) {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(erros.ErrCredenciaisInvalidas, nil))
		return
	}

	token, err := services.NewJWTService().GenerateToken(usuario.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, token))
}
