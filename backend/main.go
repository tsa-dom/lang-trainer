package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"
import "github.com/gin-gonic/contrib/static"

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.Run()
}