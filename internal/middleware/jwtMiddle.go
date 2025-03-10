package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 第一步：定义结构体
// MyClaims 定义结构体并继承jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们需要额外记录一个username和id字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成jwt
//  @param claims
//  @param mySigningKey
//  @return string
//  @return error
func GenToken(claims MyCustomClaims, mySigningKey []byte) (string, error) {
	// 使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signToken, nil

}

// ParseToken 解析jwt字符串为对象
//  @param signToken
//  @param mySigningKey
//  @return *MyCustomClaims
//  @return error
func ParseToken(signToken string, mySigningKey []byte) (*MyCustomClaims, error) {
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(signToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}

// 基于JWT的认证中间件
func JWTAuthMiddleware(mySigningKey []byte) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		// 从请求头中取出
		signToken := c.Request.Header.Get("Authorization")
		if signToken == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 1002,
				"msg":  "token为空",
			})
			c.Abort()
			return
		}

		// 校验token
		myclaims, err := ParseToken(signToken, mySigningKey)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 1003,
				"msg":  "token校验失败",
			})
			c.Abort()
			return
		}

		// 将用户的id放在到请求的上下文c上
		c.Set("userId", myclaims.Id)
		c.Set("userName", myclaims.Username)
		c.Next() // 后续的处理函数可以用过c.Get("userId")来获取当前请求的id
	}
	return fn
}
