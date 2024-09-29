package main

import (
	dbConection "api_pattern_go/api/database/conection"
	"api_pattern_go/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConection.ConnectDatabase()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
