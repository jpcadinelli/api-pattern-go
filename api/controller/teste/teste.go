package teste

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	queryConditionId = "id = ?"
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

	tx := dbConetion.DB.First(&t, queryConditionId, id)
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
	var (
		testes []models.Teste
		filtro models.TesteFiltro
	)

	if err := ginctx.ShouldBindJSON(&filtro); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	tx := dbConetion.DB

	if filtro.Nome != "" {
		tx = tx.Where("nome LIKE ?", "%"+filtro.Nome+"%")
	}

	tx.Find(&testes)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, testes))
}

func Atualizar(ginctx *gin.Context) {
	var (
		t    models.Teste
		tOld models.Teste
	)

	if err := ginctx.ShouldBindJSON(&t); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	tx := dbConetion.DB.First(&tOld, queryConditionId, t.Id)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}
	if tx.RowsAffected == 0 {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(errors.New("teste não foi encontrado"), nil))
		return
	}

	updateItems := map[string]interface{}{
		"nome": t.Nome,
	}

	tx = dbConetion.DB.Model(&tOld).Where("id = ?", t.Id).Updates(updateItems)
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, tOld))
}

func Deletar(ginctx *gin.Context) {
	id := ginctx.Param("id")

	tx := dbConetion.DB.Where(queryConditionId, id).Delete(&models.Teste{})
	if tx.Error != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(tx.Error, nil))
		return
	}
	if tx.RowsAffected == 0 {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(nil, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}
