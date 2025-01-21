package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/database"
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

    c.JOSN(http.StatusOk, gin.H{
        "message" : "Success retrieve post data",
        "posts" : posts,
    })
}
