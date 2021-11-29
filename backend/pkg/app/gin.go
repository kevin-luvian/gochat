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

type ValidationError struct {
	Msg  string `json:"message"`
	Errs []VErr `json:"errors"`
}

func (app *Gin) BindAndValid(form interface{}) (errCode int) {
	err := app.C.Bind(form)
	if err != nil {
		app.Response(http.StatusNotAcceptable, ValidationError{
			Msg:  errc.GetMsg(errc.InvalidParams),
			Errs: []VErr{},
		})
		return errc.InvalidParams
	}

	verrs := ValidateStruct(validate, form)
	if len(verrs) > 0 {
		app.Response(http.StatusNotAcceptable, ValidationError{
			Msg:  errc.GetMsg(errc.InvalidParams),
			Errs: verrs,
		})
		return errc.InvalidParams
	}

	return errc.Success
}
