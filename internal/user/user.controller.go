package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users, err := GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"users":  "Failed to retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"users":  users,
	})
}

func Get(c *gin.Context) {
	userID := c.Param("user_id")
	convUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	var user UserDto
	row, err := GetUser(&user, uint(convUserID))
	if row == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "no user found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}

func GetByQuery(c *gin.Context) {
	query := c.Query("query")

	users, err := GetUserByQuery(query)
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"users":  users,
	})
}
