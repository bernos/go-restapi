package todos

import "github.com/bernos/go-restapi/router"

var Routes = router.Routes{
	router.Route{
		"TodoIndex",
		"GET",
		"/",
		TodoIndex,
	},
	router.Route{
		"TodoShow",
		"GET",
		"/{todoId}",
		TodoShow,
	},
	router.Route{
		"TodoCreate",
		"POST",
		"/",
		TodoCreate,
	},
}
