package app

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var validatorsFn = []struct {
	Tag string
	VFn validator.Func
}{
	{
		Tag: "nestr",
		VFn: func(fl validator.FieldLevel) bool {
			return strings.TrimSpace(fl.Field().String()) != ""
		},
	},
	{
		Tag: "validurl",
		VFn: func(fl validator.FieldLevel) bool {
			fUrl := fl.Field().String()

			_, err := url.ParseRequestURI(fUrl)
			if err != nil {
				return false
			}

			u, err := url.Parse(fUrl)
			if err != nil || u.Scheme == "" || u.Host == "" {
				return false
			}

			return true
		},
	},
}

func makeValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	for _, vfn := range validatorsFn {
		err := v.RegisterValidation(vfn.Tag, vfn.VFn)
		if err != nil {
			logrus.Error(err)
		}
	}
	return v
}

func ValidateStruct(v *validator.Validate, o interface{}) []VErr {
	err := v.Struct(o)
	if err != nil {
		return parseValidationErrors(err.(validator.ValidationErrors))
	}
	return []VErr{}
}

type VErr struct {
	Field   string `json:"field" example:"input_field"`
	Message string `json:"message" example:"input_field must have a value!"`
	Value   string `json:"value" example:""`
}

func parseValidationErrors(Errors validator.ValidationErrors) []VErr {
	verrs := make([]VErr, 0, len(Errors))
	for _, err := range Errors {
		verrs = append(verrs, VErr{
			Field:   err.Field(),
			Value:   err.Param(),
			Message: err.Translate(trans),
		})
	}
	return verrs
}
