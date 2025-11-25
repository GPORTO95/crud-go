package repository

import (
	"context"
	"os"

	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init create user repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		logger.Error("Error getting user json value", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error getting user json value", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
