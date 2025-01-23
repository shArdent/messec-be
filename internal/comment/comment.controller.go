package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
)

func PostComment(c *gin.Context) {
	postId := c.Param("post_id")

	var newComment Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	value, err := strconv.ParseUint(postId, 10, 64)
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

		newComment.UserID = &convUintUserId

	}

	newComment.PostID = uint(value)

	if err := Create(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment posted",
	})
}

func DeleteComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	convCommentID, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
		return
	}

	err = Delete(&Comment{}, uint(convCommentID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Failed to delete comment",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted",
	})
}
