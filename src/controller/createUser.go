package controller

import (
	"github.com/gporto95/crud-go/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Teste")
	c.JSON(err.Code, err)
}
