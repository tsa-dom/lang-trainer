package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.Run()
}