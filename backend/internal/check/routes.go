package check

import "gochat/lib/router"

var Routes = router.Routes{
	router.Route{
		Uri:     "/ok",
		Method:  router.GET,
		Handler: healthCheck,
	},
	router.Route{
		Uri:     "/fail",
		Method:  router.GET,
		Handler: failedCheck,
	},
}
