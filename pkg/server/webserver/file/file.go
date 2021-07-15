package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/conf"
	"ledger/lib"
	"ledger/sdk"
	"path"
	"strings"
	"time"
)

func (w *FileWebServer) Upload(c *gin.Context) {
	file, err := c.FormFile("icon")
	if err != nil {
		sdk.Log.Errorf("get file error :%s", err.Error())
		w.Error(c, 5000, "get file error")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" {
		sdk.Log.Errorf("file type error:%s", fileExt)
		w.Error(c, 5000, "file type error")
		return
	}

	if _, err := lib.Cmd("mkdir", "-p", fmt.Sprintf("%s/%d", conf.Config.IconPath, u.UserId)); err != nil {
		sdk.Log.Errorf("create file path error:%s", err.Error())
		w.Error(c, 5000, "create file path error")
		return
	}

	//存文件
	t := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s_%s", t, file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("%s/%d/%s", conf.Config.IconPath, u.UserId, fileName)); err != nil {
		sdk.Log.Errorf("save file error:%s", err.Error())
		w.Error(c, 5000, "save file error")
		return
	}
	w.Success(c, nil)
}
