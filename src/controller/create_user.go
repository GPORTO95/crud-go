package controller

import (
	"fmt"

	"github.com/gporto95/crud-go/src/controller/model/request"

	"github.com/gporto95/crud-go/src/configuration/rest_err"

	"github.com/gporto95/crud-go/src/configuration/validation"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
			errRest := validation.ValidateUserError(err)

		c.JSON(restErr.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
