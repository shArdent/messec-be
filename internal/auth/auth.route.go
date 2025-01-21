package auth

import "github.com/gin-gonic/gin"

func SetupRoutes(g *gin.RouterGroup) {
	auth := g.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
    auth.GET("/logout", Logout)
}
