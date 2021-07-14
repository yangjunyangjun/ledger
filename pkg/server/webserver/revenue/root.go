package revenue

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
)

type RevenueWebServer struct {
	webbase.BaseServer
}

func (w *RevenueWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	revenueGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl), middlware.CheckLogin())
	revenueGroup.POST("add_revenue", w.addRevenue)
	revenueGroup.GET("revenue_list", w.RevenueList)
	revenueGroup.PUT("update_revenue", w.UpdateRevenue)
	//
	revenueGroup.POST("add_revenue_type", w.AddRevenueType)
	revenueGroup.GET("revenue_type_list", w.RevenueTypeList)
	revenueGroup.DELETE("del_revenue_type", w.DelRevenueType)

}
