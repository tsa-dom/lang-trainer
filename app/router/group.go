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

func addWordToGroup(c *gin.Context) {
	word := groups.Word{}
	word.OwnerId = getAuthorizedUser(c).Id

	if err := c.BindJSON(&word); err != nil {
		errorResponse(c, 400, err.Error())
		return
	}
	createdWord := groups.Word{}
	createdWord, err := groups.CreateWord(word)
	if err != nil {
		errorResponse(c, 500, err)
		return
	}

	items := createdWord.Items
	wordItems, err := groups.AddItemsToWord(createdWord.Id, items)
	if err != nil {
		errorResponse(c, 500, err)
		return
	}
	createdWord.Items = wordItems

	c.JSON(http.StatusAccepted, gin.H{
		"word": createdWord,
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
