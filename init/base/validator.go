package base

import (
	initial "blog/init"
	en2 "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	ch "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/sirupsen/logrus"
	"reflect"
)

var (
	validate   *validator.Validate
	uni        *ut.UniversalTranslator
	translator ut.Translator
)

// Validate 获取验证器
func Validate() *validator.Validate {
	return validate
}

// Uni 获取语言转换器
func Uni() *ut.UniversalTranslator {
	return uni
}

// Translate 获取翻译器
func Translate(locale string) ut.Translator {
	var (
		err   error
		found bool
	)
	translator, found = uni.GetTranslator(locale)
	if !found {
		logrus.Error("no found translator")
	}
	switch locale {
	case "en":
		err = en.RegisterDefaultTranslations(validate, translator)
		// todo 注册一次后切换语言，tag不会随语言切换
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en_comment")
		})
	case "zh_tw":
		err = zh_tw.RegisterDefaultTranslations(validate, translator)
	default:
		err = ch.RegisterDefaultTranslations(validate, translator)
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
	}
	if err != nil {
		logrus.Error(err)
	}
	return translator
}

type ValidatorStarter struct {
	initial.BaseStarter
}

func (v *ValidatorStarter) Init(ctx initial.StarterContext) {
	validate = validator.New()
	// 设置支持语言
	en := en2.New()
	ch := zh.New()
	tw := zh_Hant_TW.New()
	uni = ut.New(en, ch, tw)

}
