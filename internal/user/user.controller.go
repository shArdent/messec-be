package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
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
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request Data",
			"detail": err.Error(),
		})
		return
	}

	var existUser User
	err := GetUserByEmailOrUsername(&existUser, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "error",
			"detail": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Your Username or Password incorrect",
			"detail": err.Error(),
		})
		return
	}

	token, err := pkg.GenerateToken(existUser.Email, string(existUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error server",
			"detail": err.Error(),
		})
		return
	}

	c.SetCookie("jwt", token, int(time.Hour.Seconds()*24), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": true, "data": map[string]any{"id": existUser.ID, "email": existUser.Email, "username": existUser.Username}, "token": token, "message": "Berhasil login"})
}
