package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"red-packet/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validator() *validator.Validate {
	return validate
}

func Translate() ut.Translator {
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()
	// 创建消息国际化翻译器
	cn := zh.New()
	// universal translator
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := zh2.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logger.Errorf("register translation error: %v", err)
		}
	} else {
		logger.Error("translator not found...")
	}
}
