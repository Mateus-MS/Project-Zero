package main

import (
	"MODULE_PATH/backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":8080")
}
