package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

type user struct {
	id          int
	username    string
	priviledges string
}

func errorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
			return
		}

		c.Set("verification", user{
			id:          verification.Id,
			username:    verification.Username,
			priviledges: verification.Priviledges,
		})
		c.Next()
	}
}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
		}

		if verification.Priviledges != "admin" {
			errorResponse(c, 403, "are you admin?")
		}

		c.Next()
	}
}
