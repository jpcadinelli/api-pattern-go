package routes

import (
	"api_pattern_go/api/controller/login"
	"github.com/gin-gonic/gin"
)

const (
	logins = "/"
)

func loginRoutes(r *gin.RouterGroup) {
	r.POST(logins, login.Login)
}
