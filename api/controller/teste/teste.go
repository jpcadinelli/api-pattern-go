package teste

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	var test models.Teste
	if err := ginctx.ShouldBindJSON(&test); err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbConetion.DB.Create(&test)
	ginctx.JSON(http.StatusCreated, test)
}

func Visualizar(ginctx *gin.Context) {
	var tests []models.Teste
	dbConetion.DB.Find(&tests)
	ginctx.JSON(http.StatusOK, tests)
}
