package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	var newComment Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment posted",
	})
}
