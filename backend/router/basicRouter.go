package router

import (
	"gochat/lib/router"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type MyRouter struct{}

func MakeMyRouter() router.Router {
	return MyRouter{}
}

func (MyRouter) Handle(endpoint string, routes *router.Routes) {
	for _, route := range *routes {
		http.HandleFunc(
			combineUri(endpoint, route.Uri),
			route.Handler)
	}
}

func (MyRouter) Serve(port int) {
	logrus.Info("listening on port ", port)
	portStr := ":" + strconv.Itoa(port)
	http.ListenAndServe(portStr, nil)
}

func combineUri(endpoint string, uri string) string {
	if strings.TrimSpace(uri) == "/" {
		return endpoint
	}
	return endpoint + uri
}
