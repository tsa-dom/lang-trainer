package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	base := router.Group("/api/")

	ping(base.Group("/ping"))
	user(base.Group("/user"))

	router.Run()
}
