package routes

import (
	"github.com/gporto95/crud-go/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/by-id/:userId", controller.FindUserById)
	r.GET("/user/by-email/:userEmail", controller.FindUserByEmail)
	r.POST("/user/", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
