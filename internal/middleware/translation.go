package middleware

import (
	"blog/init/base"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const TranslatorKey = "trans"

// Translate 翻译中间件
func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		translator := base.Translate(locale)
		validate := base.Validate()

		_ = validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			if fl.Field().String() == "" {
				return true
			}
			matched, _ := regexp.Match(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, []byte(fl.Field().String()))
			return matched
		})

		_ = validate.RegisterTranslation("phone", translator, func(ut ut.Translator) error {
			return ut.Add("phone", "{0}格式错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})

		c.Set(TranslatorKey, translator)
		c.Next()
	}
}
