package Models

import (
	"errors"
	"imgHub/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func (this *Claims) GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := Claims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(config.SecretKey)
}

// ParseToken 解析JWT
func (this *Claims) ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
