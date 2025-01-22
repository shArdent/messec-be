package answer

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	answer := g.Group("/answers")
	answer.POST("/:question_id", PostAnswer)
}
