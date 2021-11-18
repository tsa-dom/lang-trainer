package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	users "github.com/tsa-dom/lang-trainer/app/models/users"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func SignNewUser(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	createdUser, err := users.CreateUser(g.User{
		Username:     user.Username,
		PasswordHash: passwordHash,
		Privileges:   user.Privileges,
	})
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	token, err := utils.CreateAuthToken(createdUser)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
		"user":  createdUser,
	})
}
