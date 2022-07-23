package v1

import (
	"math"

	"github.com/gin-gonic/gin"
	r "github.com/wannanbigpig/gin-layout/internal/pkg/response"
)

// HelloWorld hello world
func HelloWorld(c *gin.Context) {
	// str, ok := c.GetQuery("name")
	// if !ok {
	// 	str = "gin-layout"
	// }
	m := map[string]int64{"abc": math.MaxInt64}

	//this json will be used by javascript,(as json is javascrippt),but javascript can't use this MaxInt64,
	//so you should give a option to Marshal intt64 as string
	r.Success(c, m)
}
