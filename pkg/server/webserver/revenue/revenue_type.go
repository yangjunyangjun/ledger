package revenue

import (
	"github.com/gin-gonic/gin"
	"ledger/api/revenue"
	"ledger/api/revenue/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddRevenueTypeReq struct {
	TypeName string `json:"type_name"`
}

type DelRevenueTypeReq struct {
	Id int64 `json:"id"`
}

// @Title 新增收入种类
// @Author y18175612315@163.com
// @Description 新增收入种类
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddRevenueTypeReq true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/revenue/add_revenue_type [post]
func (w *RevenueWebServer) AddRevenueType(c *gin.Context) {
	var req AddRevenueTypeReq
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
	rev := db.RevenueType{
		TypeName: req.TypeName,
		UserId:   user.UserId,
	}
	if err := revenue.CreateRevenueType(rev); err != nil {
		sdk.Log.Errorf("add revenue_type err:%v", err)
		w.Error(c, 5000, "add revenue_type err")
	}
	w.Success(c, nil)
}

// @Title 收入种类列表
// @Author y18175612315@163.com
// @Description 收入种类列表
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/revenue/revenue_type_list [get]
func (w *RevenueWebServer) RevenueTypeList(c *gin.Context) {
	user := w.GetUser(c)
	if user != nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	rList, err := revenue.QueryRevenueTypeList(user.UserId)
	if err != nil {
		sdk.Log.Errorf("query revenue_type list err:%v", err)
		w.Error(c, 5000, "query revenue_type list err")
	}
	w.Success(c, rList)
}

// @Title 删除收入种类
// @Author y18175612315@163.com
// @Description 删除收入种类
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body DelRevenueTypeReq true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/revenue/del_revenue_type [delete]
func (w *RevenueWebServer) DelRevenueType(c *gin.Context) {
	var req DelRevenueTypeReq
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
	rt := db.RevenueType{
		UserId: user.UserId,
		Id:     req.Id,
	}
	if err := revenue.DeleteRevenueType(rt); err != nil {
		sdk.Log.Errorf("delete revenue_type list err:%v", err)
		w.Error(c, 5000, "delete revenue_type list err")
	}
	w.Success(c, nil)
}
