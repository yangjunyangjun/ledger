package sdk

import (
	"github.com/go-redis/redis"
	"ledger/conf"
	"os"
)

var RedisCli *redis.Client

func newRedis() {
	redisConf := conf.Config.Redis
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     redisConf.Ip,
		Password: redisConf.Password,
		DB:       0,
	})
	if _, err := RedisCli.Ping().Result(); err != nil {
		Log.Errorf("init redis error :%v", err)
		os.Exit(1)
	}
	Log.Info("init redis success !!")
}
