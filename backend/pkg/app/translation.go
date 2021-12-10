package app

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)

var VTranslations = []VTranslation{
	defaultVTranslation("required", "{0} must have a value!"),
	defaultVTranslation("passwd", "{0} password is not strong enough!"),
}

func makeTranslator(v *validator.Validate) ut.Translator {
	en := en.New()
	uni := ut.New(en, en)
	tr, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, tr)
	for _, vt := range VTranslations {
		if err := v.RegisterTranslation(vt.Tag, tr, vt.RegFn, vt.TransFn); err != nil {
			logrus.Panic("cant register translation for ", vt.Tag)
		}
	}
	for _, vtr := range validators {
		defTranslation := defaultVTranslation(vtr.Tag, vtr.Trl)
		if err := v.RegisterTranslation(vtr.Tag, tr, defTranslation.RegFn, defTranslation.TransFn); err != nil {
			logrus.Panic("cant register translation for ", vtr.Tag)
		}
	}
	return tr
}

type VTranslation struct {
	Tag     string
	RegFn   validator.RegisterTranslationsFunc
	TransFn validator.TranslationFunc
}

func defaultVTranslation(tag, msg string) VTranslation {
	return VTranslation{
		Tag: tag,
		RegFn: func(ut ut.Translator) error {
			return ut.Add(tag, msg, true)
		},
		TransFn: func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		},
	}
}
