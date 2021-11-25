package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func HttpRecorder(
	handler func(c *gin.Context),
	body io.Reader,
	authHandler gin.HandlerFunc,
	token string,
) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	if authHandler != nil {
		engine.Use(authHandler)
	}
	engine.POST("/", handler)

	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", body)
	ctx.Request.Header.Set("Authorization", "Bearer "+token)
	engine.HandleContext(ctx)
	return w
}
