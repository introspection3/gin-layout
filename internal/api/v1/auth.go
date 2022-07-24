package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/internal/error_code"
	"github.com/wannanbigpig/gin-layout/internal/middleware"
	r "github.com/wannanbigpig/gin-layout/internal/response"
	"github.com/wannanbigpig/gin-layout/internal/service"
	"github.com/wannanbigpig/gin-layout/internal/validator"
	"github.com/wannanbigpig/gin-layout/internal/validator/form"
)

// Login 一个跑通全流程的示例，业务代码未补充完整
func Login(c *gin.Context) {
	// 初始化参数结构体
	loginForm := form.LoginForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}
	// 实际业务调用
	result, err := service.Login(loginForm.UserName, loginForm.Password)
	// 根据业务返回值判断业务成功 OR 失败
	if err != nil {
		r.Fail(c, 1, err.Error())
		return
	}

	r.Ok(c, result)
}

var mySigningKey = []byte("lqzisnb")

func GetJwt(c *gin.Context) {
	var userName string = ""
	var password string = ""

	_ = password
	id := 123
	claim := middleware.MyCustomClaims{
		Username: userName,
		Id:       id,
	}
	token, err := middleware.GenToken(claim, mySigningKey)
	if err != nil {
		r.Write(c, error_code.AuthorizationError, http.StatusOK, err.Error(), 3)
		return
	}
	r.Ok(c, token)
}
