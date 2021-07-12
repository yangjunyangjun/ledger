package router

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webserver/bill"
	"ledger/pkg/server/webserver/user"
)

const (
	App     = "ledger"
	Version = "v1"
)

var (
	u = new(user.UserWebServer)
	b = new(bill.BillWebServer)
)

func Setup(g *gin.Engine) {
	u.Setup(App, Version, "user", g)
	b.Setup(App, Version, "bill", g)
}
