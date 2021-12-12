package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	groups "github.com/tsa-dom/lang-trainer/app/models/groups"
	templates "github.com/tsa-dom/lang-trainer/app/models/templates"
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

func ModifyGroup(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	group := g.Group{}
	if err := c.BindJSON(&group); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	err := groups.ModifyGroup(user.Id, group)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"group": group,
	})
}

func RemoveGroups(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	groupIds := g.GroupIds{}
	if err := c.BindJSON(&groupIds); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	err := groups.RemoveGroups(user.Id, groupIds)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"groupIds": groupIds.Ids,
	})
}

func RemoveWords(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	wordIds := g.WordIds{}
	if err := c.BindJSON(&wordIds); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	err := groups.RemoveWords(user.Id, wordIds)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"wordIds": wordIds.Ids,
	})
}

func ModifyWord(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	word := g.Word{}
	if err := c.BindJSON(&word); err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	err := groups.ModifyWord(user.Id, word)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"word": word,
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

	err = groups.AddWordToGroup(createdWord.GroupId, createdWord.Id)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"word": createdWord,
	})
}

func AddTemplate(c *gin.Context) {
	template := g.Template{}
	user := utils.GetAuthorizedUser(c)

	if err := c.BindJSON(&template); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	createdTemplate, err := templates.CreateTemplate(user.Id, template)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"template": createdTemplate,
	})
}

func ModifyTemplate(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	template := g.Template{}
	if err := c.BindJSON(&template); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	err := templates.ModifyTemplate(user.Id, template)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"template": template,
	})
}

func RemoveTemplates(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	templateIds := g.TemplateIds{}
	if err := c.BindJSON(&templateIds); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
	}

	err := templates.RemoveTemplates(user.Id, templateIds)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"templateIds": templateIds.Ids,
	})
}

func GetTemplates(c *gin.Context) {
	user := utils.GetAuthorizedUser(c)

	templates, err := templates.GetTemplates(user.Id)
	if err != nil {
		utils.ErrorResponse(c, 500, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"templates": templates,
	})
}
