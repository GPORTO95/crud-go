package repository

import (
	"os"
	"context"
	"fmt"

	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
	"github.com/gporto95/crud-go/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/gporto95/crud-go/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("no user found with email %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully", 
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByEmail"))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("no user found with id %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserById repository executed successfully", 
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByEmail"))

	return converter.ConvertEntityToDomain(*userEntity), nil
}