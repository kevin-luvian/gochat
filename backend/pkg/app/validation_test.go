package app

import (
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func vtestsetup() (*validator.Validate, ut.Translator) {
	v := makeValidator()
	return v, makeTranslator(v)
}

func TestValidUrl(t *testing.T) {
	vl, _ := vtestsetup()

	rBody := struct {
		RUrl string `validate:"validurl"`
	}{"http://localhost:3000"}

	if verrs := ValidateStruct(vl, rBody); len(verrs) > 0 {
		t.Fatal("valid url is invalid. ", verrs)
	}

	invalidUrls := []string{"", "localhost::3000"}
	for _, url := range invalidUrls {
		rBody.RUrl = url
		if verrs := ValidateStruct(vl, rBody); len(verrs) == 0 {
			t.Fatal("invalid url is valid. ", url)
		}
	}
}

func TestNotEmptyString(t *testing.T) {
	vl, _ := vtestsetup()

	rBody := struct {
		Val string `validate:"nestr"`
	}{"http://localhost:3000"}

	if verrs := ValidateStruct(vl, rBody); len(verrs) > 0 {
		t.Fatal("not empty string is invalid. ", verrs)
	}

	rBody.Val = ""
	if verrs := ValidateStruct(vl, rBody); len(verrs) == 0 {
		t.Fatal("empty string is valid. ", verrs)
	}
}
