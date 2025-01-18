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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"users":  users,
	})
}
