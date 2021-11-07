package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func errorResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func setVerification(c *gin.Context, verification utils.Claims) {
	c.Set("verification", AuthorizedUser{
		Id:          verification.Id,
		Username:    verification.Username,
		Priviledges: verification.Priviledges,
	})
}

func getAuthorizedUser(c *gin.Context) AuthorizedUser {
	user, _ := c.Get("verification")
	ctx := context.WithValue(c, authKey{}, user)
	authorizedUser := ctx.Value(authKey{}).(AuthorizedUser)
	return authorizedUser
}

func getCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	return config
}
