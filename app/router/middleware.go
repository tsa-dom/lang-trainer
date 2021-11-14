package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
			return
		}
		setVerification(c, *verification)

		c.Next()
	}
}

func AuthorizeTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
			return
		}

		privileges := verification.Privileges
		if privileges != "teacher" && privileges != "admin" {
			errorResponse(c, 403, "are you teacher")
			return
		}

		setVerification(c, *verification)

		c.Next()
	}
}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		verification, err := utils.VerifyUser(c.Request.Header.Get("Authorization"))
		if err != nil {
			errorResponse(c, 403, err.Error())
			return
		}

		if verification.Privileges != "admin" {
			errorResponse(c, 403, "are you admin?")
			return
		}

		setVerification(c, *verification)

		c.Next()
	}
}
