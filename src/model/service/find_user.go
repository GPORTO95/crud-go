package service

import (
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
	"github.com/gporto95/crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById services", zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmail(email)
}
