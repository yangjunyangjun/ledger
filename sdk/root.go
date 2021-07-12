package sdk

import "ledger/conf"

//初始化公用组件
func init() {
	//初始化log
	newLog(conf.Config.Log.FIleName)
	Log.Println("init log success")
	//初始化db
	newDB()
	Log.Println("init db success")
	//初始化redis
	newRedis()
	Log.Println("init redis success")
}
