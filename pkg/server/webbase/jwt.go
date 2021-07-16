package webbase

import (
	"github.com/dgrijalva/jwt-go"
	"ledger/conf"

	"time"
)

var jwtSecret = []byte(conf.Config.Jwt.Secret)

type Claims struct {
	UserId   int64
	Username string
	Role     int64
	Email    string
	jwt.StandardClaims
}

// 产生token的函数
func (c *Claims) GenerateToken() (string, error) {
	// 设置过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	c.ExpiresAt = expireTime.Unix()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//
	return nil, err
}
