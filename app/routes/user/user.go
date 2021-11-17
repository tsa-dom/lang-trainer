package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/models/groups"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"user": utils.GetAuthorizedUser(c),
	})
}

func GetGroups(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	groups, err := groups.GetGroups(user.Id)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"groups": groups,
	})
}

func GetWordsInGroup(c *gin.Context) {
	group := groups.Group{}
	if err := c.BindJSON(&group); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	words, err := groups.GetWordsInGroup(group.Id)
	if err != nil {
		utils.ErrorResponse(c, 500, "no words found or error in db")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"words": words,
	})

}
