package consume

import (
	"github.com/gin-gonic/gin"
	"ledger/api/consume"
	"ledger/api/consume/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddConsumeTypeReq struct {
	TypeName string `json:"type_name"`
}

type DelConsumeTypeReq struct {
	Id int64 `json:"id"`
}

// @Title 新增消费种类
// @Author y18175612315@163.com
// @Description 新增消费种类
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddConsumeTypeReq true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/add_consume_type [post]
func (w *ConsumeWebServer) AddConsumeType(c *gin.Context) {
	var req AddConsumeTypeReq
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	user := w.GetUser(c)
	if user != nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	ct := db.ConsumeType{
		TypeName: req.TypeName,
		UserId:   user.UserId,
	}
	if err := consume.CreateConsumeType(ct); err != nil {
		sdk.Log.Errorf("query cost list err:%v", err)
		w.Error(c, 5000, "query cost list err")
	}
	w.Success(c, nil)
}

// @Title 消费种类列表
// @Author y18175612315@163.com
// @Description 消费种类列表
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/consume_type_list [get]
func (w *ConsumeWebServer) ConsumeTypeList(c *gin.Context) {
	user := w.GetUser(c)
	if user != nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	cList, err := consume.ConsumeTypeList(user.UserId)
	if err != nil {
		sdk.Log.Errorf("query cost list err:%v", err)
		w.Error(c, 5000, "query cost list err")
	}
	w.Success(c, cList)
}

// @Title 删除消费种类
// @Author y18175612315@163.com
// @Description 删除种类列表
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body DelConsumeTypeReq true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/del_consume_type [delete]
func (w *ConsumeWebServer) DelConsumeType(c *gin.Context) {
	var req DelConsumeTypeReq
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	user := w.GetUser(c)
	if user != nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	ct := db.ConsumeType{
		UserId: user.UserId,
		Id:     req.Id,
	}
	if err := consume.DeleteConsumeType(ct); err != nil {
		sdk.Log.Errorf("query cost list err:%v", err)
		w.Error(c, 5000, "query cost list err")
	}
	w.Success(c, nil)
}
