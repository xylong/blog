package middleware

import (
	"blog/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
)

// Translate 翻译中间件
func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := dto.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh.RegisterDefaultTranslations(dto.Validate, trans)
		case "en":
			en.RegisterDefaultTranslations(dto.Validate, trans)
		default:
			zh.RegisterDefaultTranslations(dto.Validate, trans)
		}

		c.Set("trans", trans)
		c.Next()
	}
}
