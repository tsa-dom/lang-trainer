package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/models/words"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func group(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		groups, err := words.GetGroups(verification.Id)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"groups": groups,
		})
	})

	route.POST("/", func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		group := words.Group{}
		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		group.OwnerId = verification.Id
		createdGroup, err := words.CreateGroup(group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"group": createdGroup,
		})
	})
}
