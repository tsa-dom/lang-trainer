package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/language-trainer/app/utils"
)

type CreateUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Priviledges string `json:"priviledges"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func user(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {

	})

	route.POST("/signup/", func(c *gin.Context) {
		user := CreateUser{}

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusAccepted, &user)
	})

	route.POST("/login/", func(c *gin.Context) {
		user := LoginUser{}

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		token, err := utils.GetAuthToken(user.Username, user.Password)
		if err != nil {
			c.AbortWithError(http.StatusNonAuthoritativeInfo, err)
			return
		}

		c.JSON(http.StatusAccepted, token)
	})

	route.PUT("/", func(c *gin.Context) {

	})

	route.DELETE("/", func(c *gin.Context) {

	})
}
