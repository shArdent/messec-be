package middlewares

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

func CORSMiddlewareLib() cors.Config {
	originString := viper.GetString("ALLOW_ORIGINS")
	originArray := strings.Split(originString, " ")
	return cors.Config{
		AllowOrigins:     originArray,
		AllowMethods:     []string{"PUT", "GET", "OPTIONS", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
}
