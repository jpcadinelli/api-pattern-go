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
	err := dbConection.RunMigrations()
	if err != nil {
		return
	}
}

func iniciaRotasAPI() {
	router := gin.Default()
	router = routes.SetupRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
