package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	teacher "github.com/tsa-dom/lang-trainer/app/routes/teacher"
)

func Run() {
	apiGateway := gin.Default()

	corsConfig := getCorsConfig()
	apiGateway.Use(cors.New(corsConfig))

	apiGateway.Use(static.Serve("/", static.LocalFile("./build", true)))
	apiGateway.Static("/my/", "./build")

	api := apiGateway.Group("/api/")

	apiAdmin := api.Group("/admin/")
	apiAdmin.Use(AuthorizeAdmin())
	apiAdmin.GET("/user/", getUser)
	apiAdmin.POST("/signup/", signNewUser)

	apiTeacher := api.Group("/teacher/")
	apiTeacher.Use(AuthorizeTeacher())
	apiTeacher.POST("/groups/", teacher.AddGroup)
	apiTeacher.POST("/word/", teacher.AddWordToGroup)

	apiPrivate := api.Group("/my/")
	apiPrivate.Use(AuthorizeUser())
	apiPrivate.GET("/", getUser)
	apiPrivate.GET("/groups/", getGroups)

	apiPrivate.POST("/words/", getWordsInGroup)

	api.POST("/login/", loginUser)

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	apiGateway.Run()

}
