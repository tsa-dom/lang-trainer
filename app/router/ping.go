package router

import "github.com/gin-gonic/gin"

func ping(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
}
