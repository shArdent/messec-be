package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/pkg"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := pkg.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
