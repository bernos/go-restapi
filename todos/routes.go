package todos

import (
	"github.com/bernos/go-restapi/router"
)

var c = &TodoController{}

var Routes = router.Routes{
	router.Route{
		"TodoIndex",
		"GET",
		"/",
		c.Action(c.TodoIndex),
	},
	router.Route{
		"TodoShow",
		"GET",
		"/{todoId}",
		c.Action(c.TodoShow),
	},
	router.Route{
		"TodoCreate",
		"POST",
		"/",
		c.Action(c.TodoCreate),
	},
}
