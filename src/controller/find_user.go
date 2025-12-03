package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/gporto95/crud-go/src/configuration/logger"
	"github.com/gporto95/crud-go/src/configuration/rest_err"
	"github.com/gporto95/crud-go/src/model"
	"github.com/gporto95/crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	logger.Info("Init findUserById controller", zap.String("journey", "findUserById"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserById"))

		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	logger.Info("Result from service", zap.Bool("has_error", err != nil), zap.String("journey", "findUserById"))
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err, zap.String("journey", "findUserById"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully", zap.String("journey", "findUserById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "findUserByEmail"))

		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services", err, zap.String("journey", "findUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserLogged(c *gin.Context) {
	logger.Info("Init findUserLogged controller", zap.String("journey", "findUserLogged"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userDomain, err := uc.service.FindUserByIDServices(user.GetID())
	if err != nil {
		logger.Error("Error trying to call findUserLogged services", err, zap.String("journey", "findUserLogged"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("findUserLogged controller executed successfully", zap.String("journey", "findUserLogged"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
