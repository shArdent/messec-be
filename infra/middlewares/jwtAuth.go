package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		err := pkg.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Not authenticated",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
