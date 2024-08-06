package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Translators = ut.New(en.New())

var NewTranslator = ut.New

type (
	Validate = validator.Validate
	Register = func(v *validator.Validate, trans ut.Translator) (err error)
)

// ApplyTranslator 添加翻译 默认 "en"
//
//	import (
//		"github.com/go-playground/locales/zh"
//		"github.com/tsingshaner/gin/validator"
//	)
//
//	func main() {
//	    // 覆盖默认
//		_ = validator.ApplyTranslator(zh.New(), true, zhTranslator.RegisterDefaultTranslations)
//	}
func ApplyTranslator(t locales.Translator, override bool, register Register) error {
	if err := Translators.AddTranslator(zh.New(), true); err != nil {
		return err
	}

	if ginValidator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if t, ok := Translators.GetTranslator("zh"); ok {
			return register(ginValidator, t)
		}
	}

	return nil
}
