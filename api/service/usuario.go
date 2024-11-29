package service

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"api_pattern_go/api/repository"
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	BearerSchema = "Bearer "
)

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
