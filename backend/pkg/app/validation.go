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
	for _, vfn := range validatorsFn {
		err := v.RegisterValidation(vfn.Tag, vfn.VFn)
		if err != nil {
			logrus.Error(err)
		}
	}
	return v
}

type VErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
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

type vValidator struct {
	Tag string
	VFn validator.Func
}

var validatorsFn = []vValidator{
	{
		Tag: "nestr",
		VFn: func(fl validator.FieldLevel) bool {
			return strings.TrimSpace(fl.Field().String()) != ""
		},
	},
}
