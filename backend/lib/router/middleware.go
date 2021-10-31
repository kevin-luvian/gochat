package router

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type MiddlewareHandlerFunc func(http.ResponseWriter, *http.Request) bool

func CreateMiddleware(logic MiddlewareHandlerFunc) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			if ok := logic(res, req); ok {
				next(res, req)
			}
		}
	}
}
