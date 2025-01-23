package post

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	posts := g.Group("/posts")
	posts.GET("/:user_id", GetAllPostByUserId)
	posts.POST("/", middlewares.JwtAuthMiddleware(), CreateNewPost)
	posts.DELETE("/:post_id", middlewares.JwtAuthMiddleware(), DeletePost)
}
