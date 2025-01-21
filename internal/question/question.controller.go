package question

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	userID := c.Param("user_id")
	var newQuestion Question

	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Error Bad Request",
			"detail": err.Error(),
		})
		return
	}

	convUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid UserID",
			"detail": err.Error(),
		})
		return
	}

	newQuestion.UserID = uint(convUserID)
	if err := Create(&newQuestion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error creating question",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "question created",
	})
}

func GetAllQuestionByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	questions, err := GetQuestionByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Error getting questions data",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Success retrieve post data",
		"questions": questions,
	})
}
