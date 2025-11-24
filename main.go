package main

import (
	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/controller/routes"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to start application... ", err)
	}
}
