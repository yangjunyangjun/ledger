package user

import (
	"github.com/gin-gonic/gin"
	"ledger/api/captcha"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type Verify struct {
	Id   string `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
}


// @Title 获取图片验证码
// @Author y18175612315@163.com
// @Description 获取图片验证码
// @Tags 用户相关接口
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/captcha [get]
func (w *UserWebServer) Captcha(c *gin.Context) {
	id, str := captcha.GetCaptcha()
	rst := map[string]string{
		"id":     id,
		"string": str,
	}
	w.Success(c, rst)
}



// @Title 图片验证码校验
// @Author y18175612315@163.com
// @Description 图片验证码校验
// @Tags 用户相关接口
// @Param body body	Verify true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/captcha_verify [post]
func (w *UserWebServer) Verify(c *gin.Context) {
	var req Verify
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	if !captcha.Verify(req.Id, req.Code) {
		sdk.Log.Error("Verify error")
		w.Error(c, 5000, "Verify error")
		return
	}
	w.Success(c, nil)
}
