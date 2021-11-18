package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	groups "github.com/tsa-dom/lang-trainer/app/models/groups"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func AddGroup(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	group := g.Group{}
	if err := c.BindJSON(&group); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	group.OwnerId = user.Id
	createdGroup, err := groups.CreateGroup(group)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"group": createdGroup,
	})
}

func AddWordToGroup(c *gin.Context) {
	word := g.Word{}
	word.OwnerId = utils.GetAuthorizedUser(c).Id

	if err := c.BindJSON(&word); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	createdWord, err := groups.CreateWord(word)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}

	items := createdWord.Items
	wordItems, err := groups.AddItemsToWord(createdWord.Id, items)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}
	createdWord.Items = wordItems

	err = groups.AddWordToGroup(createdWord.GroupId, createdWord.Id)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"word": createdWord,
	})
}
