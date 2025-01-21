package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPostByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	posts, err := GetPostsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Error getting posts data",
			"detail": err.Error(),
		})
		return
	}

    c.JSON(http.StatusOK, gin.H{
        "message" : "Success retrieve post data",
        "posts" : posts,
    })
}
