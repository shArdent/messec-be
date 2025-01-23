package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	comment := g.Group("/comments")
	// create comment by postId
	comment.POST("/:post_id", PostComment)
	comment.DELETE("/:comment_id", middlewares.JwtAuthMiddleware(), DeleteComment)
}
