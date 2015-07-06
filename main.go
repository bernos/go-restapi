package main

import (
	"log"

	"github.com/bernos/go-restapi/application"
	"github.com/bernos/go-restapi/config"
	"github.com/bernos/go-restapi/todos"
	"github.com/unrolled/render"
)

func main() {

	config := config.NewApiConfiguration()

	configureRender(config)
	configureModules(config)

	app := application.NewApplication()
	app.Use(Logger)

	log.Fatal(app.ListenAndServe(":8080", NewRouter()))
}

func configureRender(c *config.ApiConfiguration) {
	c.Render = render.New()
}

func configureModules(c *config.ApiConfiguration) {
	todos.Configure(c)
}
