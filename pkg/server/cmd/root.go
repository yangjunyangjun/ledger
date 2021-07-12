package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ledger/docs"
	"ledger/lib"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/router"
	"ledger/sdk"
	"os"
)

func Run() {
	g := gin.New()
	// 添加log日志中间件
	g.Use(middlware.LoggerToFile(),
	)
	// 注册路由
	router.Setup(g)
	// 获取本机ip
	ip, err := lib.GetIp()
	if err != nil {
		sdk.Log.Error("get ip error")
		panic("get ip error")
	}

	// 初始化 添加swgger 文件
	out := lib.Cmd("swag", "init")
	sdk.Log.Println(string(out))

	//添加swgger 路由
	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	url := ginSwagger.URL(fmt.Sprintf("http://%s:8080/swagger/doc.json", ip)) // The url pointing to API definition
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 启动服务
	if err := g.Run(fmt.Sprintf("%s:8080", ip)); err != nil {
		os.Exit(1)
	}
}
