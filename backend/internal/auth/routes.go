package auth

import "gochat/lib/router"

var Routes = router.Routes{
	router.Route{
		Uri:     "/login",
		Method:  router.GET,
		Handler: temp,
	},
	router.Route{
		Uri:     "/login/google",
		Method:  router.GET,
		Handler: loginGoogle,
	},
	router.Route{
		Uri:     "/google",
		Method:  router.GET,
		Handler: authGoogle,
	},
}
