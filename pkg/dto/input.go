package dto

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func init() {
	// 设置支持语言
	en := en2.New()
	zh := zh2.New()
	zh_tw := zh_Hant_TW.New()
	// 设置国际化翻译器
	Uni = ut.New(en, zh, zh_tw)
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
