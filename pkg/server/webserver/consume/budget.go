package consume

import (
	"github.com/gin-gonic/gin"
	"ledger/api/consume"
	"ledger/api/consume/db"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
	"time"
)

type AddBudgetReq struct {
	Money float64 `json:"money" form:"money" binding:"required"` // 预算金额
}

type QueryBudgetReq struct {
	YearMon string `json:"year_mon" form:"year_mon" binding:"required"`  //年月
}

// @Title 新增预算
// @Author y18175612315@163.com
// @Description 新增预算
// @Tags 预算相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body AddBudgetReq true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/add_budget [post]
func (w *ConsumeWebServer) AddBudget(c *gin.Context) {
	var req AddBudgetReq
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	budget := db.Budget{
		UserId:  u.UserId,
		Money:   req.Money,
		YearMon: time.Now().Format("2006-01"),
	}
	if err := consume.AddBudget(budget); err != nil {
		sdk.Log.Error("create budget error:", err)
		w.Error(c, 5000, "create budget error")
		return
	}
	w.Success(c, nil)
}

// @Title 查询预算详情
// @Author y18175612315@163.com
// @Description 查询预算详情
// @Tags 预算相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param year_mon query string true "年月"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/consume/budget_details [get]
func (w *ConsumeWebServer) BudgetDetails(c *gin.Context) {
	var req QueryBudgetReq
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u := w.GetUser(c)
	if u == nil {
		sdk.Log.Error("user empty")
		w.Error(c, 5000, "user empty")
		return
	}
	budgetMoney, consumeMoney, err := consume.QueryBudget(u.UserId, req.YearMon)
	if err != nil {
		sdk.Log.Error("query budget err:", err)
		w.Error(c, 5000, "query budget err")
		return
	}
	w.Success(c, map[string]float64{"budget_money": budgetMoney, "consume_money": consumeMoney})
}
