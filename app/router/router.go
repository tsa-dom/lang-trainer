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
	router.Static("/my/", "./build")

	api := router.Group("/api/")

	ping(api.Group("/ping"))
	user(api.Group("/user"))
	word(api.Group("/word"))
	group(api.Group(("/group")))

	router.Run()
}
