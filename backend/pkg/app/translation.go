package app

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var VTranslations = []VTranslation{
	defaultVTranslation("required", "{0} must have a value!"),
	defaultVTranslation("nestr", "{0} must not be an empty string!"),
	defaultVTranslation("passwd", "{0} password is not strong enough!"),
	defaultVTranslation("validurl", "{0} must be a valid url"),
}

func makeTranslator(v *validator.Validate) ut.Translator {
	en := en.New()
	uni := ut.New(en, en)
	tr, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, tr)
	for _, vt := range VTranslations {
		_ = v.RegisterTranslation(vt.Tag, tr, vt.RegFn, vt.TransFn)
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
