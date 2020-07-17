package dto

import (
	"blog/init/base"
	"github.com/gin-gonic/gin"
)

// Param 参数
type Param interface {
	// 校验数据
	Check(*gin.Context, interface{}) error
}

// Input 输入
type Input struct {
}

// Check 数据校验
func (input *Input) Check(c *gin.Context, param interface{}) error {
	if err := c.ShouldBind(param); err != nil {
		return err
	}
	return base.Validate().Struct(param)
}
