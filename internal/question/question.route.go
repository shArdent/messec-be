package question

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	question := g.Group("/questions")
	question.GET("/user_id", GetAllQuestionByUserId)
	question.POST("/:user_id", CreateQuestion)
}
