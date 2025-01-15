package router

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	v1group := r.Group("/api/v1")

	user.SetupRoutes(&user.UserController{DB: "apa ajah"}, v1group)
	
}
