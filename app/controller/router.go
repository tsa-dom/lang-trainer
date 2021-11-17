package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	admin "github.com/tsa-dom/lang-trainer/app/routes/admin"
	teacher "github.com/tsa-dom/lang-trainer/app/routes/teacher"
	user "github.com/tsa-dom/lang-trainer/app/routes/unauthorized"
	private "github.com/tsa-dom/lang-trainer/app/routes/user"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func Run() {
	apiGateway := gin.Default()

	corsConfig := utils.GetCorsConfig()
	apiGateway.Use(cors.New(corsConfig))

	apiGateway.Use(static.Serve("/", static.LocalFile("./build", true)))
	apiGateway.Static("/my/", "./build")

	api := apiGateway.Group("/api/")

	apiAdmin := api.Group("/admin/")
	apiAdmin.Use(AuthorizeAdmin())
	apiAdmin.POST("/signup/", admin.SignNewUser)

	apiTeacher := api.Group("/teacher/")
	apiTeacher.Use(AuthorizeTeacher())
	apiTeacher.POST("/groups/", teacher.AddGroup)
	apiTeacher.POST("/word/", teacher.AddWordToGroup)

	apiPrivate := api.Group("/my/")
	apiPrivate.Use(AuthorizeUser())
	apiPrivate.GET("/", private.GetUser)
	apiPrivate.GET("/groups/", private.GetGroups)
	apiPrivate.POST("/words/", private.GetWordsInGroup)

	api.POST("/login/", user.Login)

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	apiGateway.Run()

}
