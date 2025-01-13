package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/routes"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	routes.SetupRoutes(r)

    r.Run(":8080")
}
