package dto

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func init() {
	en := en2.New()
	zh := zh2.New()
	Uni = ut.New(en, zh)
	Validate = validator.New()
}

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
	return Validate.Struct(param)
}
