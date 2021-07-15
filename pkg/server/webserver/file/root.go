package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
)

type FileWebServer struct {
	webbase.BaseServer
}

func (w *FileWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	fileGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl), middlware.CheckLogin())
	fileGroup.POST("upload", w.Upload)
}
