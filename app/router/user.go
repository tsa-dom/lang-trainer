package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/models/users"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func getUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"user": getAuthorizedUser(c),
	})
}

func loginUser(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		errorResponse(c, 400, err.Error())
		return
	}

	authUser, err := users.GetUserByUsername(user.Username)
	if err != nil {
		errorResponse(c, 400, err.Error())
		return
	}
	check := utils.CheckPasswordHash(user.Password, authUser.PasswordHash)

	if !check {
		errorResponse(c, 400, "password and hash not match")
		return
	}

	token, err := utils.CreateAuthToken(authUser.Username)
	if err != nil {
		errorResponse(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token":    token,
		"username": authUser.Username,
	})
}
