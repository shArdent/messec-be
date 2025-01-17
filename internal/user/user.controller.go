package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Register(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request Data",
			"detail": err.Error(),
		})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to create new user",
			"detail": err.Error(),
		})
		return
	}

	newUser.Password = string(hashed)
	if err := CreateNew(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to create new user",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func Login(c *gin.Context) {
	var user *User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request Data",
			"detail": err.Error(),
		})
		return
	}
}
