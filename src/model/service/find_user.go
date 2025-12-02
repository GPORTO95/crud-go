package service

import (
	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById services", zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword services", zap.String("journey", "findUserByEmailAndPassword"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
