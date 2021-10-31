package router

type Router interface {
	Handle(endpoint string, routes *Routes)
	Serve(port int)
}
