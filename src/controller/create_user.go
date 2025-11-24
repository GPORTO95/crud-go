package controller

import (
	"net/http"
	"github.com/gporto95/crud-go/src/controller/model/request"
	"github.com/gporto95/crud-go/src/controller/model/response"
	"github.com/gporto95/crud-go/src/configuration/validation"
	"github.com/gin-gonic/gin"
	"github.com/gporto95/crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "123",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
