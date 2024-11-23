package routes

import (
	"api_pattern_go/api/controller/login"
	"github.com/gin-gonic/gin"
)

func loginRoutes(r *gin.RouterGroup) {
	r.POST(route, login.Login)
}
