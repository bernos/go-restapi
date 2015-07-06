package todos

import (
	"github.com/bernos/go-restapi/config"
	"github.com/bernos/go-restapi/router"
)

var (
	Routes router.Routes
)

func ConfigureRoutes(config *config.ApiConfiguration) {
	c := &TodoController{render: config.Render}

	Routes = router.Routes{
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
}
