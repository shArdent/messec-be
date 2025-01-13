package user

import "github.com/gin-gonic/gin"

func SetupRoutes(uc *UserController, g *gin.RouterGroup) {
	user := g.Group("/users")

	user.GET("/", uc.GetAll)
}
