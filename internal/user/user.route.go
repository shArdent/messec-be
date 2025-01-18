package user

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	user := g.Group("/users")
	user.GET("/", GetAll)
}
