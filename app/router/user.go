package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/models/users"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Priviledges string `json:"priviledges"`
}

func getUser(c *gin.Context) {

	user, _ := c.Get("verification")
	log.Println(user)

	c.JSON(http.StatusAccepted, gin.H{
		"username":    "user",
		"priviledges": "user",
	})
}

func signNewUser(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := users.CreateUser(users.User{
		Username:     user.Username,
		PasswordHash: passwordHash,
		Priviledges:  user.Priviledges,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.CreateAuthToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token":    token,
		"username": user.Username,
	})
}

func loginUser(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	authUser, err := users.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	check := utils.CheckPasswordHash(user.Password, authUser.PasswordHash)

	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password and hash not match",
		})
		return
	}

	token, err := utils.CreateAuthToken(authUser.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token":    token,
		"username": authUser.Username,
	})
}
