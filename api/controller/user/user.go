package user

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	var u models.User

	if err := ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u.Password = services.SHA256Encoder(u.Password)

	tx := dbConetion.DB.Create(&u)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, u))
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
