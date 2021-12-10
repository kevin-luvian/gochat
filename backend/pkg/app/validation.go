package app

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func makeValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	for _, vfn := range validators {
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
}

func parseValidationErrors(Errors validator.ValidationErrors) []VErr {
	verrs := make([]VErr, 0, len(Errors))
	for _, err := range Errors {
		verrs = append(verrs, VErr{
			Field:   err.Field(),
			Message: err.Translate(trans),
		})
	}
	return verrs
}
