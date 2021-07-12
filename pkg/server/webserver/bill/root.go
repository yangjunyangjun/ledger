package bill

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
)

type BillWebServer struct {
	webbase.BaseServer
}

func (w *BillWebServer) Setup(app, version, baseUrl string, g *gin.Engine) {
	billGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl), middlware.CheckLogin())
	billGroup.POST("add_bill", w.addBill)
	billGroup.GET("bill_list", w.BillList)

}
