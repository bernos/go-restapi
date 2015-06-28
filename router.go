package main

import (
	"github.com/bernos/restapi/router"
	"github.com/bernos/restapi/todos"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	registerSubroutes(router.PathPrefix("/todos").Subrouter(), todos.Routes)

	return router
}

func registerSubroutes(router *mux.Router, routes router.Routes) {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
