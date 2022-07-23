package routers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/assets"
	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/internal/middleware"
	"github.com/wannanbigpig/gin-layout/internal/pkg/error_code"
	response2 "github.com/wannanbigpig/gin-layout/internal/pkg/response"
)

func SetRouters() *gin.Engine {
	r := gin.New()
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))
	r.StaticFS("assets", http.FS(assets.Static))
	fav := func(c *gin.Context) {
		if c.Request.RequestURI != "/favicon.ico" {
			return
		}
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			status := http.StatusOK
			if c.Request.Method != "OPTIONS" {
				status = http.StatusMethodNotAllowed
			}
			c.Header("Allow", "GET,HEAD,OPTIONS")
			c.AbortWithStatus(status)
			return
		}

		c.Header("Content-Type", "image/x-icon")
		data, err := assets.Static.ReadFile("static/favicon.ico")
		if err == nil {
			c.Writer.Write(data)
		}

	}
	// 初始化默认静态资源
	r.Use(fav)

	// 设置模板资源

	r.Use(
		middleware.RequestCostHandler(),
		middleware.CustomRecovery(),
		middleware.CorsHandler(),
	)

	if config.Config.Debug == false {
		// 生产模式
		ReleaseRouter()
		r.Use(
			middleware.CustomLogger(),
		)
	} else {
		// 开发调试模式
		r.Use(
			gin.Logger(),
		)
	}

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// 设置 API 路由
	setApiRoute(r)

	r.NoRoute(func(c *gin.Context) {
		response2.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, error_code.NotFound)
	})

	return r
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

}
