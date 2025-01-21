package post

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	posts := g.Group("/posts")
	posts.GET("/:user_id", GetAllPostByUserId)
    posts.POST("/:user_id", CreateNewPost)
}
