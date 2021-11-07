package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/models/groups"
)

func getGroups(c *gin.Context) {
	user := getAuthorizedUser(c)

	groups, err := groups.GetGroups(user.Id)
	if err != nil {
		errorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"groups": groups,
	})
}

func addGroup(c *gin.Context) {
	user := getAuthorizedUser(c)

	group := groups.Group{}
	if err := c.BindJSON(&group); err != nil {
		errorResponse(c, 400, err.Error())
		return
	}

	group.OwnerId = user.Id
	createdGroup, err := groups.CreateGroup(group)
	if err != nil {
		errorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"group": createdGroup,
	})
}

func getWordsInGroup(c *gin.Context) {
	group := groups.Group{}
	if err := c.BindJSON(&group); err != nil {
		errorResponse(c, 400, err.Error())
		return
	}

	words, err := groups.GetWordsInGroup(group.Id)
	if err != nil {
		errorResponse(c, 500, "no words found or error in db")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"words": words,
	})

}
