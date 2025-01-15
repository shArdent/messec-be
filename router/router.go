package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes() *gin.Engine {
	environtment := viper.GetBool("DEBUG")
	if environtment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	RegisterRoutes(router)

	return router
}
