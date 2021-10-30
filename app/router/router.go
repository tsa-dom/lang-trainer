package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	base := router.Group("/api/")

	ping(base.Group("/ping"))
	user(base.Group("/user"))

	router.Run()
}
