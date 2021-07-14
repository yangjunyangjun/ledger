package revenue

import (
	"github.com/gin-gonic/gin"
	"ledger/api/revenue"
	"ledger/api/revenue/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddRevenueRequest struct {
	Money  float64 `json:"money"`
	Type   int64   `json:"type"`
	Remark string  `json:"remark"`
}

type RevenueListRequest struct {
	Type      int64  `json:"type" form:"type"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	webbase.RequestPage
}

type UpdateRevenueRequest struct {
	Id int64 `json:"id"`
	AddRevenueRequest
}

// @Title 新增收入记录
// @Author y18175612315@163.com
// @Description 新增收入记录
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddRevenueRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/revenue/add_revenue [post]
func (w *RevenueWebServer) addRevenue(c *gin.Context) {
	var req AddRevenueRequest
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
	rev := db.Revenue{
		UserId: u.UserId,
		Money:  req.Money,
		Type:   req.Type,
		Remark: req.Remark,
	}
	if err := revenue.CreateRevenue(rev); err != nil {
		sdk.Log.Errorf("add revenue error:%v", err)
		w.Error(c, 5000, "add revenue error")
		return
	}
	w.Success(c, nil)
}

// @Title 收入记录列表
// @Author y18175612315@163.com
// @Description 收入记录列表
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param query query RevenueListRequest true "param数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/revenue/revenue_list [get]
func (w *RevenueWebServer) RevenueList(c *gin.Context) {
	var req RevenueListRequest
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
	cList, err := revenue.QueryRevenue(u.UserId, req.StartTime, req.EndTime, req.Type, req.Offset, req.Limit)
	if err != nil {
		sdk.Log.Errorf("query revenue error:%v", err)
		w.Error(c, 5000, "query revenue error")
		return
	}
	w.Success(c, cList)
}

// @Title 更新收入记录
// @Author y18175612315@163.com
// @Description 更新收入记录
// @Tags 收入管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body UpdateRevenueRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/update_revenue [put]
func (w *RevenueWebServer) UpdateRevenue(c *gin.Context) {
	var req UpdateRevenueRequest
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
	rev := db.Revenue{
		UserId: u.UserId,
		Money:  req.Money,
		Type:   req.Type,
		Remark: req.Remark,
	}
	if err := revenue.UpdateRevenue(rev); err != nil {
		sdk.Log.Errorf("update revenue err:", err)
		w.Error(c, 5000, "update revenue err")
		return
	}
	w.Success(c, nil)
}
