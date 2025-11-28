package service

import (
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
	"github.com/gporto95/crud-go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUserServices(string) *rest_err.RestErr
}
