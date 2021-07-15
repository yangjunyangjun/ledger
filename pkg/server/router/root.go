package router

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webserver/consume"
	"ledger/pkg/server/webserver/file"
	"ledger/pkg/server/webserver/revenue"
	"ledger/pkg/server/webserver/user"
)

const (
	App     = "ledger"
	Version = "v1"
)

var (
	userServer    = new(user.UserWebServer)
	consumeServer = new(consume.ConsumeWebServer)
	revenueServer = new(revenue.RevenueWebServer)
	fileServer    = new(file.FileWebServer)
)

func Setup(g *gin.Engine) {
	userServer.Setup(App, Version, "user", g)
	consumeServer.Setup(App, Version, "consume", g)
	revenueServer.Setup(App, Version, "revenue", g)
	fileServer.Setup(App, Version, "file", g)
}
