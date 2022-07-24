package boot

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/data"
	"github.com/wannanbigpig/gin-layout/internal/routers"
	"github.com/wannanbigpig/gin-layout/internal/validator"
	"github.com/wannanbigpig/gin-layout/pkg/logger"
)

func init() {
	var configPath string

	flag.StringVar(&configPath, "c", "", "请输入配置文件绝对路径")
	flag.Parse()

	// 1、初始化配置
	config.InitConfig(configPath)

	// 2、初始化zap日志
	logger.InitLogger()

	// 3、初始化数据库
	data.InitData()

	// 4、初始化验证器
	validator.InitValidatorTrans("zh")
}

func Run() {
	r := routers.SetRouters()
	addr := fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		// service connections
		var err error = nil
		logger.Logger.Info("server listen:" + addr)
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Sugar().Errorf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Warn("shutdown server,please wait 5s ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Warn("server shutdown:" + err.Error())
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Logger.Warn("Server exiting")
}
