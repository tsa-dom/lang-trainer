package utils

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

type authKey struct{}

func ErrorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func SetVerification(c *gin.Context, verification g.Claims) {
	c.Set("verification", g.AuthorizedUser{
		Id:         verification.Id,
		Username:   verification.Username,
		Privileges: verification.Privileges,
	})
}

func GetAuthorizedUser(c *gin.Context) g.AuthorizedUser {
	user, _ := c.Get("verification")
	ctx := context.WithValue(c, authKey{}, user)
	authorizedUser := ctx.Value(authKey{}).(g.AuthorizedUser)
	return authorizedUser
}

func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	return config
}
