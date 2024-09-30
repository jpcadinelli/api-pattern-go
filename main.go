package main

import (
	dbConection "api_pattern_go/api/database/conection"
	"api_pattern_go/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	carregaDadosIniciais()
	iniciaConfigBanco()

	iniciaRotasAPI()
}

func carregaDadosIniciais() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func iniciaConfigBanco() {
	dbConection.ConnectDatabase()
	dbConection.MakeMigrations()
}

func iniciaRotasAPI() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
