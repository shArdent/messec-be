package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	answer := g.Group("/answers")
	answer.Use(middlewares.JwtAuthMiddleware())

	answer.POST("/:question_id", PostAnswer)
	answer.DELETE("/:answer_id", DeleteAnswer)
}
