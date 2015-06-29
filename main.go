package main

import (
	"log"

	"github.com/bernos/go-restapi/application"
)

func main() {

	app := application.NewApplication()
	app.Use(Logger)

	log.Fatal(app.ListenAndServe(":8080", NewRouter()))
}
