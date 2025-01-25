package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/middlewares"
	"github.com/spf13/viper"
)

func SetupRoutes() *gin.Engine {
	environtment := viper.GetBool("DEBUG")
	if environtment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(middlewares.CORSMiddlewareLib()))

	RegisterRoutes(router)

	return router
}
