package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/internal/user"
	"github.com/shardent/messec-be/pkg"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var newUser user.User

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
	var existUser user.User
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request Data",
			"detail": err.Error(),
		})
		return
	}

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

	convUserId := strconv.FormatUint(uint64(existUser.ID), 10)

	token, err := pkg.GenerateToken(existUser.Email, convUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error server",
			"detail": err.Error(),
		})
		return
	}

	c.SetCookie("access_token", token, int(time.Hour.Seconds()*24), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": true, "data": map[string]any{"id": existUser.ID, "email": existUser.Email, "username": existUser.Username}, "token": token, "message": "Berhasil login"})
}

func Logout(c *gin.Context) {
	_, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
