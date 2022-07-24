package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/internal/error_code"
)

func Resp() *Response {
	// 初始化response
	return &Response{
		HttpCode: http.StatusOK,
		ErrCode:  0,
		Message:  "",
		Data:     nil,
		Cost:     "",
	}
}

// Success 业务成功响应
func Success(c *gin.Context, data ...any) {
	if data != nil {
		Resp().WithDataSuccess(c, data[0])
		return
	}
	r := Resp()
	r.SetErrCode(error_code.SUCCESS)
	r.Ok = true
	r.Json(c)
}

// Fail 业务失败响应
func Fail(c *gin.Context, code int, data ...any) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code)
		return
	}
	Resp().FailCode(c, code)
}

type Response struct {
	HttpCode int    `json:"httpCode"`
	Ok       bool   `json:"ok"`
	ErrCode  int    `json:"errCode"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Cost     string `json:"cost"`
}

// Fail 错误返回
func (r *Response) Fail(c *gin.Context) {
	r.SetErrCode(error_code.FAILURE)
	r.Json(c)
}

// FailCode 自定义错误码返回
func (r *Response) FailCode(c *gin.Context, code int, msg ...string) {
	r.SetErrCode(code)
	r.Ok = false
	if msg != nil {
		r.WithMessage(msg[0])
	}
	r.Json(c)
}

// WithDataSuccess 成功后需要返回值
func (r *Response) WithDataSuccess(c *gin.Context, data interface{}) {
	r.SetErrCode(error_code.SUCCESS)
	r.Ok = true
	r.WithData(data)
	r.Json(c)
}

// SetErrCode 设置返回code码
func (r *Response) SetErrCode(code int) *Response {
	r.ErrCode = code
	r.Ok = code == 1
	return r
}

// SetHttpCode 设置http状态码
func (r *Response) SetHttpCode(code int) *Response {
	r.HttpCode = code
	return r
}

// WithData 设置返回data数据
func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

// WithMessage 设置返回自定义错误消息
func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

var errorText = &error_code.ErrorText{
	Language: config.Config.Language,
}

// Json 返回 gin 框架的 HandlerFunc
func (r *Response) Json(c *gin.Context) {
	if r.Message == "" {
		r.Message = errorText.Text(r.ErrCode)
	}

	r.Cost = time.Since(c.GetTime("requestStartTime")).String()
	c.JSON(200, *r)
}
