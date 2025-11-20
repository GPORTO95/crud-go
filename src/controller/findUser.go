package controller

import "github.com/gin-gonic/gin"

func FindUserById(c *gin.Context) {
	userId := c.Param("userId")
	c.JSON(200, gin.H{
		"message": "Find user by ID: " + userId,
	})
}

func FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	c.JSON(200, gin.H{
		"message": "Find user by ID: " + userEmail,
	})
}
