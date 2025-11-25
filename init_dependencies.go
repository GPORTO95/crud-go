package main

import (
	"github.com/gporto95/crud-go/src/controller"
	"github.com/gporto95/crud-go/src/model/repository"
	"github.com/gporto95/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
