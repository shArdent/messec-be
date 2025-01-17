package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users, err := GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"users":  "Failed to retrieve users",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"users":  users,
	})
}

func Create(c *gin.Context) {
	var newUser *User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request Data",
			"detail": err.Error(),
		})
	}

	if err := CreateNew(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to create new user",
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}
