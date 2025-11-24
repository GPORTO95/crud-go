package main

import (
	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/controller/routes"
	"github.com/gporto95/crud-go/src/controller"
	"github.com/gporto95/crud-go/src/model/service"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to start application... ", err)
	}
}
