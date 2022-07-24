package routers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/assets"
	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/internal/error_code"
	"github.com/wannanbigpig/gin-layout/internal/middleware"
	"github.com/wannanbigpig/gin-layout/internal/response"
)

func SetRouters() *gin.Engine {
	eng := gin.New()
	eng.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))
	eng.StaticFS("assets", http.FS(assets.Static))
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
	eng.Use(fav)

	// 设置模板资源

	eng.Use(
		middleware.RequestCostHandler(),
		middleware.CustomRecovery(),
		middleware.CorsHandler(),
	)

	if !config.Config.Debug {
		// 生产模式
		ReleaseRouter()
		eng.Use(
			middleware.CustomLogger(),
		)
	} else {
		// 开发调试模式
		eng.Use(
			gin.Logger(),
		)
	}

	// ping
	eng.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// 设置 API 路由
	setApiRoute(eng)

	eng.NoRoute(func(c *gin.Context) {
		response.FailHttpStatus(c, error_code.NotFound, http.StatusNotFound, "")
	})

	return eng
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard
}
