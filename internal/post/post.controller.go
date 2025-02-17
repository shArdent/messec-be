package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
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
		"message": "Success retrieve post data",
		"posts":   posts,
	})
}

func CreateNewPost(c *gin.Context) {
	userId, err := pkg.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Error unauthorized",
			"detail": err.Error(),
		})
		return
	}

	var post Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Error getting posts data",
			"detail": err.Error(),
		})
		return
	}

	value, err := strconv.ParseUint(userId.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error server",
			"detail": err.Error(),
		})
		return
	}

	post.UserID = uint(value)

	if err = CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error server",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success create new post",
	})
}

func DeletePost(c *gin.Context) {
	postID := c.Param("post_id")
	convPostID, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	err = Delete(&Post{}, uint(convPostID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to delete post",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
