package utils

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type AuthorizedUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Privileges string `json:"privileges"`
	Token      string `json:"token"`
}

type authKey struct{}

func ErrorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func SetVerification(c *gin.Context, verification Claims) {
	c.Set("verification", AuthorizedUser{
		Id:         verification.Id,
		Username:   verification.Username,
		Privileges: verification.Privileges,
	})
}

func GetAuthorizedUser(c *gin.Context) AuthorizedUser {
	user, _ := c.Get("verification")
	ctx := context.WithValue(c, authKey{}, user)
	authorizedUser := ctx.Value(authKey{}).(AuthorizedUser)
	return authorizedUser
}

func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	return config
}
