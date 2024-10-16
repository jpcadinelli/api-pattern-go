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
	dbConetion.DB.Create(&t)
	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, t))
}

func Visualizar(ginctx *gin.Context) {
	var t models.Teste
	id := ginctx.Param("id")
	dbConetion.DB.First(&t, "id = ?", id)
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, t))
}

func Listar(ginctx *gin.Context) {
	var testes []models.Teste
	dbConetion.DB.Find(&testes)
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, testes))
}
