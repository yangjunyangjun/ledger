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
	consumeGroup := g.Group(fmt.Sprintf("%s/%s/%s", app, version, baseUrl), middlware.CheckLogin())
	consumeGroup.POST("add_consume", w.addConsume)
	consumeGroup.GET("consume_list", w.ConsumeList)
	consumeGroup.PUT("update_consume", w.UpdateConsume)

	consumeGroup.POST("add_consume_type", w.AddConsumeType)
	consumeGroup.GET("consume_type_list", w.ConsumeTypeList)
	consumeGroup.DELETE("del_consume_type", w.DelConsumeType)

}

