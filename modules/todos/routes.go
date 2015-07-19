package todos

import (
	"github.com/bernos/go-restapi/application"
	"github.com/bernos/go-restapi/config"
)

var (
	Routes application.Routes
)

func ConfigureRoutes(config *config.Configuration) {
	c := &TodoController{render: config.Render}

	Routes = application.Routes{
		application.Route{
			"TodoIndex",
			"GET",
			"/",
			c.Action(c.TodoIndex),
		},
		application.Route{
			"TodoShow",
			"GET",
			"/{todoId}",
			c.Action(c.TodoShow),
		},
		application.Route{
			"TodoCreate",
			"POST",
			"/",
			c.Action(c.TodoCreate),
		},
	}
}
