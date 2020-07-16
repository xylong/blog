package middleware

import (
	"blog/internal/dto"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-playground/validator/v10/translations/zh_tw"
	"reflect"
	"regexp"
)

const TranslatorKey = "trans"

// Translate 翻译中间件
func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := dto.Uni.GetTranslator(locale)

		switch locale {
		case "en":
			en.RegisterDefaultTranslations(dto.Validate, trans)
			dto.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
		case "zh_tw":
			zh_tw.RegisterDefaultTranslations(dto.Validate, trans)
		default:
			zh.RegisterDefaultTranslations(dto.Validate, trans)
			dto.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})
		}

		dto.Validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			if fl.Field().String() == "" {
				return true
			}
			matched, _ := regexp.Match(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, []byte(fl.Field().String()))
			return matched
		})

		dto.Validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "{0}格式错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})

		c.Set(TranslatorKey, trans)
		c.Next()
	}
}
