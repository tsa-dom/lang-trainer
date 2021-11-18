package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	users "github.com/tsa-dom/lang-trainer/app/models/users"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func Login(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	authUser, err := users.GetUserByUsername(user.Username)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	check := utils.CheckPasswordHash(user.Password, authUser.PasswordHash)

	if !check {
		utils.ErrorResponse(c, 400, "password and hash not match")
		return
	}

	token, err := utils.CreateAuthToken(authUser)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token":    token,
		"username": authUser.Username,
	})
}
