package app

import (
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

/*
validators list:

- nestr: not empty string

- validurl: validate url syntax
*/
var validators = []struct {
	Tag string
	Trl string
	VFn validator.Func
}{
	{
		Tag: "nestr",
		Trl: "{0} must not be an empty string",
		VFn: func(fl validator.FieldLevel) bool {
			return strings.TrimSpace(fl.Field().String()) != ""
		},
	},
	{
		Tag: "validurl",
		Trl: "{0} must be a valid url",
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
