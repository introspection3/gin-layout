package v1

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

//https://gin-gonic.com/zh-cn/docs/examples/
//https://github.com/eddycjy/go-gin-example/blob/master/routers/api/v1/article.go
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func HelloWorld(g *gin.Context) {

	g.JSON(http.StatusOK, gin.H{
		"id": math.MaxInt64,
	})
}
