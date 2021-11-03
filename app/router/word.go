package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func word(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"username":    verification.Username,
			"priviledges": verification.Priviledges,
		})
	})
}
