package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
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

func Update(c *gin.Context) {
	userID, err := pkg.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Error request",
			"detail": err.Error(),
		})
		return
	}
	var updatedData UserDto
	var existUser User

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	convUserID, err := strconv.ParseUint(userID.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	userCount, err := GetUser(&existUser, uint(convUserID))
	if userCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "no user found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	updatedData.ID = uint(convUserID)

	if err := UpdateUser(&existUser, updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error server",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User data updated",
		"user":    updatedData,
	})
}
