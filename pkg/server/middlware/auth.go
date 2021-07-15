package middlware

import (
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
	"time"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		var u *webbase.Claims

		//从redis获取用户信息
		if err := sdk.RedisCli.Get(auth).Scan(&u); err != nil {
			//redis不存在解析token
			u, err = webbase.ParseToken(auth)
			//解析失败代表token过期
			if err != nil {
				c.JSON(401, "API token required")
				c.Abort()
				return
			}
		}
		// redis有用户信息 或 token解析成功，将用户信息set到redis
		if err := sdk.RedisCli.Set(auth, u, time.Hour*24*1).Err(); err != nil {
			sdk.Log.Errorf("redis set error %v ", err)
			c.JSON(5000, "redis set err")
			return
		}
		c.Set("user", u)
	}
}
