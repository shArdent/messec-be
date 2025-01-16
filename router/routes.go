package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	v1group := r.Group("/api/v1")

	user.SetupRoutes(&user.UserController{DB: "apa ajah"}, v1group)
	
}
