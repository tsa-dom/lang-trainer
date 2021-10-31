package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/db"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Priviledges string `json:"priviledges"`
}

func user(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"username":    verification.Username,
			"priviledges": verification.Priviledges,
		})
	})

	route.POST("/signup/", func(c *gin.Context) {
		user := User{}

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		passwordHash, err := utils.HashPassword(user.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id := db.CreateUser(user.Username, passwordHash, user.Priviledges)
		if id < 1 {
			c.AbortWithError(http.StatusNotAcceptable, errors.New("database error, username already exists"))
			return
		}

		token, err := utils.CreateAuthToken(user.Username, user.Priviledges)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"token":    token,
			"username": user.Username,
		})
	})

	route.POST("/login/", func(c *gin.Context) {
		user := User{}

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		authInfo, err := db.UserAuthInfo(user.Username)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		check := utils.CheckPasswordHash(user.Password, authInfo.PasswordHash)

		if !check {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		token, err := utils.CreateAuthToken(authInfo.Username, authInfo.Priviledges)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"token":    token,
			"username": authInfo.Username,
		})
	})

	route.PUT("/", func(c *gin.Context) {

	})

	route.DELETE("/", func(c *gin.Context) {

	})
}
