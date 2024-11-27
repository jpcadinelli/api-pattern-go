package global

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIdUsuarioLogado(ginctx *gin.Context) (uuid.UUID, error) {
	var (
		id  uuid.UUID
		err error
	)
	const BearerSchema = "Bearer "

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
