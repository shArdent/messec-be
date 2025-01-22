package answer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
)

func PostAnswer(c *gin.Context) {
	questionId := c.Param("question_id")

	var newAnswer Answer

	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	value, err := strconv.ParseUint(questionId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	userId, err := pkg.ExtractTokenId(c)
	if err == nil {
		conv64UserId, _ := strconv.ParseUint(userId.(string), 10, 64)

		convUintUserId := uint(conv64UserId)

		newAnswer.UserID = convUintUserId

	}

	newAnswer.QuestionID = uint(value)

	if err := CreateComment(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Answer Posted posted",
	})
}
