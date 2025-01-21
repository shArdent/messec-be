package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/internal/auth"
	"github.com/shardent/messec-be/internal/comment"
	"github.com/shardent/messec-be/internal/post"
	"github.com/shardent/messec-be/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	v1group := r.Group("/api/v1")

	user.SetupRoutes(v1group)
	auth.SetupRoutes(v1group)
	post.SetupRoutes(v1group)
	comment.SetupRoutes(v1group)
}
