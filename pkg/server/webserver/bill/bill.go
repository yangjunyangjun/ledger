package bill

import (
	"github.com/gin-gonic/gin"
	"ledger/api/bill"
	"ledger/api/bill/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddRequest struct {
	Amount float64 `json:"amount"`
	Type   int64   `json:"type"`
	Remark string  `json:"remark"`
}

type BillListRequest struct {
	Type      int64  `json:"type" form:"type"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	webbase.RequestPage
}

// @Title 新增消费记录
// @Author y18175612315@163.com
// @Description 新增消费记录
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/bill/add_bill [post]
func (w *BillWebServer) addBill(c *gin.Context) {
	var req AddRequest
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
	b := db.Bill{
		UserId: u.UserId,
		Amount: req.Amount,
		Type:   req.Type,
		Remark: req.Remark,
	}
	if err := bill.AddBill(b); err != nil {
		sdk.Log.Errorf("add bill error:%v", err)
		w.Error(c, 5000, "add bill error")
		return
	}
	w.Success(c, nil)
}

// @Title 消费记录列表
// @Author y18175612315@163.com
// @Description 消费记录列表
// @Tags 消费管理相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param query query BillListRequest true "param数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/bill/bill_list [get]
func (w *BillWebServer) BillList(c *gin.Context) {
	var req BillListRequest
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
	bList, err := bill.QueryByBill(req.StartTime, req.EndTime, u.UserId, req.Type, req.Limit, req.Offset)
	if err != nil {
		sdk.Log.Errorf("query bill error:%v", err)
		w.Error(c, 5000, "query bill error")
		return
	}
	w.Success(c, bList)
}
