package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
)

type UserWebServer struct {
	webbase.BaseServer
}

func (w *UserWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	userGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl))
	userGroup.POST("/login", w.login)
	userGroup.POST("/register", w.register)
	userGroup.GET("/send_activate", w.sendAct)
	userGroup.GET("/user_list", middlware.CheckLogin(), w.userList)
}
