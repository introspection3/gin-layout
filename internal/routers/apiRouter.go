package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/wannanbigpig/gin-layout/docs"
	c "github.com/wannanbigpig/gin-layout/internal/api/v1"
)

func setApiRoute(r *gin.Engine) {
	// version 1
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", c.Login)
		v1.GET("/hello-world", c.HelloWorld)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
