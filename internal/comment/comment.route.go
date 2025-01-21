package comment

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	comment := g.Group("/comments")
    // create comment by postId  
	comment.POST("/:post_id", PostComment)
}
