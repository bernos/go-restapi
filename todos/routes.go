package todos

import (
	"github.com/bernos/go-restapi/router"
)

var todoController = &TodoController{}

var Routes = router.Routes{
	router.Route{
		"TodoIndex",
		"GET",
		"/",
		todoController.Action(TodoIndex),
	},
	router.Route{
		"TodoShow",
		"GET",
		"/{todoId}",
		todoController.Action(TodoShow),
	},
	router.Route{
		"TodoCreate",
		"POST",
		"/",
		todoController.Action(TodoCreate),
	},
}
