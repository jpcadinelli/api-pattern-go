package middleware

import (
	"api_pattern_go/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		const BearerSchema = "Bearer "
		header := ginctx.Request.Header.Get("Authorization")
		if header == "" {
			ginctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}

		token := header[len(BearerSchema):]

		if !services.NewJWTService().ValidateToken(token) {
			ginctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
	}
}
