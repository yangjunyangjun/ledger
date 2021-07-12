package webbase

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
)

type WebRouter interface {
	Setup(app, version, baseUrl string, g *gin.Engine)
}

func (w *BaseServer) GetUser(c *gin.Context) *middlware.Claims {
	user, ok := c.Get("user")
	if !ok {
		return nil
	}
	u, ok := user.(*middlware.Claims)
	if !ok {
		return nil
	}
	return u
}
