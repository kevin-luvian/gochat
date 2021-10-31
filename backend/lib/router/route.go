package router

import "net/http"

type httpMethod string

// Declare typed constants each with type of status
const (
	GET    httpMethod = "GET"
	PUT    httpMethod = "PUT"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
)

type Route struct {
	Uri         string
	Method      httpMethod
	Middlewares []Middleware
	Handler     http.HandlerFunc
}

func (r Route) GetOnionHandler() http.HandlerFunc {
	var handlerFunc = r.Handler
	for i := range r.Middlewares {
		handlerFunc = r.Middlewares[len(r.Middlewares)-i-1](handlerFunc)
	}
	return handlerFunc
}

type Routes []Route
