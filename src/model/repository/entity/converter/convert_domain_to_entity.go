package converter

import (
	"github.com/gporto95/crud-go/src/model"
	"github.com/gporto95/crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
