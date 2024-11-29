package usuario

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/services"
	"api_pattern_go/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	BearerSchema = "Bearer "
)

func Criar(ginctx *gin.Context) {
	var u models.Usuario

	if err := ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u.Password = services.SHA256Encoder(u.Password)

	if err := repository.NewUsuarioRepository(dbConetion.DB).Create(&u); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, userResponse))
}

func Visualizar(ginctx *gin.Context) {
	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	u, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(id, "Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
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

func GetIdUsuarioLogado(ginctx *gin.Context) (uuid.UUID, error) {
	var (
		id  uuid.UUID
		err error
	)

	header := ginctx.Request.Header.Get("Authorization")
	if header == "" {
		return id, erros.ErrTokenInexistente
	}

	token := header[len(BearerSchema):]

	if id, err = services.NewJWTService().GetUserId(token); err != nil {
		return id, err
	}

	return id, nil
}

func GetUsuarioLogado(ginctx *gin.Context) (*models.UsuarioDTOResponse, error) {
	header := ginctx.Request.Header.Get("Authorization")
	if header == "" {
		return nil, erros.ErrTokenInexistente
	}

	token := header[len(BearerSchema):]

	id, err := services.NewJWTService().GetUserId(token)
	if err != nil {
		return nil, err
	}

	usuario, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(id, "Permissoes")
	if err != nil {
		return nil, err
	}

	userResponse := usuario.UsuarioToDTOResponse()
	return userResponse, nil
}
