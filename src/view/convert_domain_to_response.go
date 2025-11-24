package view

import (
	"github.com/gporto95/crud-go/src/controller/model/response"
	"github.com/gporto95/crud-go/src/model"
)

func ConvertDomainToResponse(domain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: domain.GetEmail(),
		Name:  domain.GetName(),
		Age:   domain.GetAge(),
	}
}