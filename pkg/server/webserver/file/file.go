package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/conf"
	"ledger/lib"
	"ledger/sdk"
	"os"
	"path"
	"strings"
	"time"
)

// @Title 上传图片
// @Author y18175612315@163.com
// @Description 上传图片
// @Tags 文件接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param icon formData file true "文件"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/file/upload [post]
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

	if _, err := lib.Cmd("mkdir", "-p", fmt.Sprintf("%s/%d", conf.Config.IconPath, u.Username)); err != nil {
		sdk.Log.Errorf("create file path error:%s", err.Error())
		w.Error(c, 5000, "create file path error")
		return
	}

	//存文件
	t := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s/%s_%s", u.Username, t, file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", conf.Config.IconPath, fileName)); err != nil {
		sdk.Log.Errorf("save file error:%s", err.Error())
		w.Error(c, 5000, "save file error")
		return
	}
	w.Success(c, lib.UrlBese64Encode(fileName))
}

// @Title 下载图片
// @Author y18175612315@163.com
// @Description 下载图片
// @Tags 文件接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param icon query string true "图片连接"
// @Router	/ledger/v1/file/download [get]
func (w *FileWebServer) Download(c *gin.Context) {
	icon := c.Param("icon")
	path := lib.UrlBese64Decode(icon)
	f, err := os.Open(fmt.Sprintf("%s/%s", conf.Config.IconPath, path))
	if err != nil {
		sdk.Log.Errorf("open file error :%s", err.Error())
		c.Status(404)
		return
	}
	defer f.Close()

	//文件流
	var content []byte
	if _,err:=f.Read(content);err!=nil{
		sdk.Log.Error("read file error: ", err)
		w.Error(c, 5000, "read file error")
		return
	}

	fileNameList := strings.Split(path, "/")
	fileName := fileNameList[len(fileNameList)-1]

	attachment := fmt.Sprintf("attachment; filename=%s", fileName)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", attachment)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	//回写到web 流媒体 形成下载
	if _, err := c.Writer.Write(content); err != nil {
		sdk.Log.Error("Error write back: ", err)
		w.Error(c, 5000, "error download")
		return
	}

}
