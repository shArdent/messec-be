package user

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	auth := g.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)

	user := g.Group("/users")
	user.GET("/", GetAll)
}
