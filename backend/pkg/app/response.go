package app

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (app *Gin) Response(httpCode, errCode int, data interface{}) {
	app.C.JSON(httpCode, Response{httpCode, "ok", data})
}
