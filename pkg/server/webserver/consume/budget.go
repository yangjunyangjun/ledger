package consume

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
)

type AddBudgetReq struct {
	UserId  int64  `json:"user_id" form:"user_id" binding:"required"`
	YearMon string `json:"user_id" form:"user_id" binding:"required"`
}

func (w *ConsumeWebServer) AddBudget(c *gin.Context) {
	var req AddBudgetReq
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
}
