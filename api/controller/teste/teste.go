package teste

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	var user models.Teste
	if err := ginctx.ShouldBindJSON(&user); err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbConetion.DB.Create(&user)
	ginctx.JSON(http.StatusCreated, user)
}

func Visualizar(ginctx *gin.Context) {
	var users []models.Teste
	dbConetion.DB.Find(&users)
	ginctx.JSON(http.StatusOK, users)
}
