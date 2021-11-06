package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func errorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
			return
		}

		c.Next()
	}
}
