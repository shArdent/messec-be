package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	answer := g.Group("/answers")
	answer.POST("/:question_id", middlewares.JwtAuthMiddleware(), PostAnswer)
}
