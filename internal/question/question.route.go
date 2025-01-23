package question

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
)

func SetupRoutes(g *gin.RouterGroup) {
	question := g.Group("/questions")
    question.GET("/:user_id", GetAllQuestionByUserId)
	question.POST("/:user_id", CreateQuestion)
    question.DELETE("/:question_id", middlewares.JwtAuthMiddleware(), DeleteQuestion)
}
