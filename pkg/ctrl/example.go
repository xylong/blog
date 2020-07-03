package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Example = &ExampleController{}

type ExampleController struct {
}

// @Summary ping
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"pong"}"
// @Router /ping [get]
func (e *ExampleController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "pong",
	})
}
