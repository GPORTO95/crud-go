package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gporto95/crud-go/src/configuration/database/mongodb"
	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
