package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	user := g.Group("/users")
	user.GET("/", GetByQuery)
	user.GET("/:user_id", Get)
    user.PUT("/", middlewares.JwtAuthMiddleware(), Update)
}
