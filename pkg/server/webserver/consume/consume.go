package consume

import (
	"github.com/gin-gonic/gin"
	"ledger/api/consume"
	"ledger/api/consume/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddConsumeRequest struct {
	Money  float64 `json:"money"`
	TypeId int64   `json:"type_id"`
	Remark string  `json:"remark"`
}

type ConsumeListRequest struct {
	TypeId    int64  `json:"type_id" form:"type_id"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	webbase.RequestPage
}

type UpdateConsumeRequest struct {
	Id int64 `json:"id"`
	AddConsumeRequest
}

// @Title 新增消费记录
// @Author y18175612315@163.com
// @Description 新增消费记录
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddConsumeRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/add_consume [post]
func (w *ConsumeWebServer) addConsume(c *gin.Context) {
	var req AddConsumeRequest
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user no permission")
		w.Error(c, 5000, "user no permission")
		return
	}
	con := db.Consume{
		UserId: u.UserId,
		Money:  req.Money,
		TypeId: req.TypeId,
		Remark: req.Remark,
	}
	if err := consume.AddConsume(con); err != nil {
		sdk.Log.Errorf("add consume error:%v", err)
		w.Error(c, 5000, "add consume error")
		return
	}
	w.Success(c, nil)
}

// @Title 消费记录列表
// @Author y18175612315@163.com
// @Description 消费记录列表
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param query query ConsumeListRequest true "param数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/consume_list [get]
func (w *ConsumeWebServer) ConsumeList(c *gin.Context) {
	var req ConsumeListRequest
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user no permission")
		w.Error(c, 5000, "user no permission")
		return
	}
	cList, err := consume.QueryByConsume(req.StartTime, req.EndTime, u.UserId, req.TypeId, req.Limit, req.Offset)
	if err != nil {
		sdk.Log.Errorf("query consume error:%v", err)
		w.Error(c, 5000, "query consume error")
		return
	}
	w.Success(c, cList)
}

// @Title 更新消费记录
// @Author y18175612315@163.com
// @Description 更新消费记录
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body UpdateConsumeRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/update_consume [put]
func (w *ConsumeWebServer) UpdateConsume(c *gin.Context) {
	var req UpdateConsumeRequest
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user no permission")
		w.Error(c, 5000, "user no permission")
		return
	}
	con := db.Consume{
		Id:     req.Id,
		UserId: u.UserId,
		Money:  req.Money,
		TypeId: req.TypeId,
		Remark: req.Remark,
	}
	if err := consume.UpdateConsume(con); err != nil {
		sdk.Log.Errorf("update consume err:%s", err.Error())
		w.Error(c, 5000, "update consume err")
		return
	}
	w.Success(c, nil)
}
