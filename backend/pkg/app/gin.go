package app

import (
	"gochat/pkg/errc"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = makeValidator()
	trans = makeTranslator(validate)
}

type Gin struct {
	C *gin.Context
}

type ErrResponse struct {
	Msg string `json:"msg"`
}

func (app *Gin) ErrResponse(httpCode, errcCode int) {
	app.C.JSON(httpCode, ErrResponse{errc.GetMsg(errcCode)})
}

func (app *Gin) OkResponse(data interface{}) {
	app.C.JSON(http.StatusOK, data)
}

func (app *Gin) Response(httpCode int, data interface{}) {
	app.C.JSON(httpCode, data)
}

func (app *Gin) BindAndValid(form interface{}) (errCode int, vErr []VErr) {
	err := app.C.Bind(form)
	if err != nil {
		return errc.InvalidParams, []VErr{}
	}

	err = validate.Struct(form)
	if err != nil {
		verrs := parseValidationErrors(err.(validator.ValidationErrors))
		return errc.FailedValidation, verrs
	}

	return errc.Success, []VErr{}
}
