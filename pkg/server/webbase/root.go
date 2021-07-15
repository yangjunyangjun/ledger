package webbase

import (
	"github.com/gin-gonic/gin"
)

type WebRouter interface {
	Setup(app, version, baseUrl string, g *gin.Engine)
}

func (w *BaseServer) GetUser(c *gin.Context) *Claims {
	user, ok := c.Get("user")
	if !ok {
		return nil
	}
	u, ok := user.(*Claims)
	if !ok {
		return nil
	}
	return u
}
