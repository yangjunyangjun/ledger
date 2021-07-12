package sdk

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"ledger/conf"
	"os"
)

var DB *gorm.DB
var err error

func newDB() {
	dbConfig := conf.Config.DataBase
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Ip, dbConfig.Db))
	if err != nil {
		Log.Errorf("connect mysql err： %v", err)
		os.Exit(1)
	}
	Log.Info("init mysql success !!")

	// 开启 Logger, 以展示详细的日志
	DB.LogMode(true)
}
