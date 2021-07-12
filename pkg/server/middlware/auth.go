package middlware

import (
	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		claim, err := ParseToken(auth)
		if err != nil {
			c.JSON(401, "API token required")
			c.Abort()
		}
		c.Set("user", claim)
	}
}
