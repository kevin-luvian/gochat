package auth

import "gochat/lib/router"

var Routes = router.Routes{
	router.Route{
		Uri:     "/login",
		Method:  router.GET,
		Handler: temp,
	},
}
