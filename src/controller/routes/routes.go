package routes

import (
	"github.com/gporto95/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/user/by-id/:userId", userController.FindUserById)
	r.GET("/user/by-email/:userEmail", userController.FindUserByEmail)
	r.POST("/user/", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
