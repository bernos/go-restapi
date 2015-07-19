package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bernos/go-restapi/application"
	"github.com/bernos/go-restapi/config"
	"github.com/bernos/go-restapi/middleware"
	"github.com/bernos/go-restapi/modules/todos"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	c := config.NewConfiguration()
	stack := middleware.New()

	configureLogging(c, stack)
	configureRender(c, stack)
	configureModules(c, stack)
	configureRouting(c, stack)

	log.Fatal(http.ListenAndServe(":8080", stack))
}

func configureLogging(c *config.Configuration, stack *middleware.MiddlewareStack) {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Set log formatting")

	stack.Use(middleware.MiddlewareFunc(LoggerMiddleware))
}

func configureRender(c *config.Configuration, stack *middleware.MiddlewareStack) {
	c.Render = render.New()
}

func configureModules(c *config.Configuration, stack *middleware.MiddlewareStack) {
	todos.Configure(c)
}

func configureRouting(c *config.Configuration, stack *middleware.MiddlewareStack) {
	c.Router = mux.NewRouter()
	c.Router.StrictSlash(true)

	registerSubroutes(c.Router.PathPrefix("/todos").Subrouter(), todos.Routes)

	stack.UseHandler(c.Router)
}

func registerSubroutes(router *mux.Router, routes application.Routes) {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}
}
