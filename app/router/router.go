package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	apiGateway := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	apiGateway.Use(cors.New(config))

	apiGateway.Use(static.Serve("/", static.LocalFile("./build", true)))
	apiGateway.Static("/my/", "./build")

	api := apiGateway.Group("/api/")

	api.Use(AuthorizeUser())

	apiAdmin := api.Group("/admin")
	apiAdmin.Use(AuthorizeAdmin())
	apiAdmin.GET("/user", getUser)

	apiPrivate := api.Group("/user")
	apiPrivate.GET("/", getUser)

	api.POST("/user", signNewUser)
	api.POST("/login", loginUser)

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGateway.Run()

}

/* func Run() {
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
} */
