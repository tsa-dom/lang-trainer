package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/language-trainer/app/utils"
)

func word(route *gin.RouterGroup) {
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
}
