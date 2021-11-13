package auth

import "gochat/lib/router"

var Endpoint = "/auth"

var Routes = router.Routes{
	router.Route{
		Uri:     "/login",
		Method:  router.GET,
		Handler: temp,
	},
	router.Route{
		Uri:     "/login/google/current",
		Method:  router.GET,
		Handler: loginGoogle,
	},
	router.Route{
		Uri:     "/google",
		Method:  router.GET,
		Handler: authGoogle,
	},
	router.Route{
		Uri:     "/refresh",
		Method:  router.GET,
		Handler: refresh,
	},
	router.Route{
		Uri:     "/tok",
		Method:  router.GET,
		Handler: parseToken,
	},
	router.Route{
		Uri:     "/login/google",
		Method:  router.POST,
		Handler: loginGoogle,
	},
}
