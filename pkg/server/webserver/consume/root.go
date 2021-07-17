package consume

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
)

type ConsumeWebServer struct {
	webbase.BaseServer
}

func (w *ConsumeWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	consumeGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl), middlware.CheckLogin())
	//消费支出
	consumeGroup.POST("add_consume", w.addConsume)
	consumeGroup.GET("consume_list", w.ConsumeList)
	consumeGroup.PUT("update_consume", w.UpdateConsume)

	//消费种类
	consumeGroup.POST("add_consume_type", w.AddConsumeType)
	consumeGroup.GET("consume_type_list", w.ConsumeTypeList)
	consumeGroup.DELETE("del_consume_type", w.DelConsumeType)

	//预算
	consumeGroup.POST("add_budget", w.AddBudget)
	consumeGroup.GET("budget_details", w.BudgetDetails)



	consumeGroup.GET("month_view", w.ConsumeView)
	consumeGroup.GET("day_view", w.ConsumeDayView)

}
