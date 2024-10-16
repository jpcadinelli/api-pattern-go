package teste

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	var t models.Teste

	if err := ginctx.ShouldBindJSON(&t); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	tx := dbConetion.DB.Create(&t)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, t))
}

func Visualizar(ginctx *gin.Context) {
	var t models.Teste

	id := ginctx.Param("id")

	tx := dbConetion.DB.First(&t, "id = ?", id)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}
	if tx.RowsAffected == 0 {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(nil, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, t))
}

func Listar(ginctx *gin.Context) {
	var testes []models.Teste

	tx := dbConetion.DB.Find(&testes)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, testes))
}
