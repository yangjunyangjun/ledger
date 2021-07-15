package middlware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
	"time"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		var (
			u   *webbase.Claims
			err error
		)

		//从redis获取用户信息
		userByte, err := sdk.RedisCli.Get(auth).Bytes()
		if err != nil || userByte == nil {
			//redis不存在解析token
			u, err = webbase.ParseToken(auth)
			//解析失败代表token过期
			if err != nil {
				c.JSON(401, "API token required")
				c.Abort()
				return
			}
			//将用户信息转为byte
			userByte, err = json.Marshal(u)
			if err != nil {
				sdk.Log.Errorf("Marshal  error %v ", err)
				c.JSON(401, "API token required")
				c.Abort()
				return
			}

		} else {
			if err = json.Unmarshal(userByte, &u); err != nil {
				sdk.Log.Errorf("UmMarshal  error %v ", err)
				c.JSON(401, "API token required")
				c.Abort()
				return

			}
		}

		// redis有用户信息 或 token解析成功，将用户信息set到redis
		if err := sdk.RedisCli.Set(auth, userByte, time.Hour*24*1).Err(); err != nil {
			sdk.Log.Errorf("redis set error %v ", err)
			c.JSON(401, "redis set err")
			c.Abort()
			return
		}
		c.Set("user", u)
	}
}
