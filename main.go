package main

import (
	"log"
	"net/http"

	"github.com/bernos/restapi/application"
)

func main() {
	router := NewRouter()
	app := application.NewApplication()

	app.
		UseMiddleware(Logger).
		UseRouter(router)

	log.Fatal(http.ListenAndServe(":8080", app))
}
