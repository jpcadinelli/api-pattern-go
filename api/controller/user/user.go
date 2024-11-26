package user

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	queryConditionId = "id = ?"
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

	userResponse := u.UserToDTOResponse()
	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, userResponse))
}

func Visualizar(ginctx *gin.Context) {
	var u models.User

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	tx := dbConetion.DB.First(&u, queryConditionId, id)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}
	if tx.RowsAffected == 0 {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(nil, nil))
		return
	}

	userResponse := u.UserToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, userResponse))
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
