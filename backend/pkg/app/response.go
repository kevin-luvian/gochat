package app

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"ok"`
	Data interface{} `json:"data" example:{}`
}

func (app *Gin) Response(httpCode, errCode int, data interface{}) {
	app.C.JSON(httpCode, Response{httpCode, "ok", data})
}
