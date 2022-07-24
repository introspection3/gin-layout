package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/internal/error_code"
)

type Result struct {
	HttpStatus int    `json:"httpStatus"`
	Ok         bool   `json:"ok"`
	ErrCode    int    `json:"errCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// Ok 业务成功响应
func Ok(c *gin.Context, data any) {
	Write(c, error_code.SUCCESS, http.StatusOK, "", data)
}

func Write[T any](c *gin.Context, errCode int, httpStatus int, msg string, data T) {
	result := &Result{
		HttpStatus: httpStatus,
		ErrCode:    errCode,
		Message:    msg,
		Data:       data,
		Ok:         errCode == 0,
	}
	c.JSON(200, result)
}

func Fail(c *gin.Context, errCode int, msg string) {
	Write(c, http.StatusOK, errCode, msg, "")
}

func FailHttpStatus(c *gin.Context, errCode int, httpStatus int, msg string) {
	Write(c, httpStatus, errCode, msg, "")
}

var errorText = &error_code.ErrorText{
	Language: config.Config.Language,
}
