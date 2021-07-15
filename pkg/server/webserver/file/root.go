package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webbase"
)

type FileWebServer struct {
	webbase.BaseServer
}

func (w *FileWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	fileGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl))
	fileGroup.POST("upload", w.Upload)
	fileGroup.GET("download", w.Download)
}
