package router

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webserver/consume"
	"ledger/pkg/server/webserver/user"
)

const (
	App     = "ledger"
	Version = "v1"
)

var (
	userServer = new(user.UserWebServer)
	consumeServer = new(consume.ConsumeWebServer)
)

func Setup(g *gin.Engine) {
	userServer.Setup(App, Version, "user", g)
	consumeServer.Setup(App, Version, "consume", g)
}
