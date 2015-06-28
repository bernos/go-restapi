package application

import "net/http"

type Application interface {
	UseRouter(http.Handler) Application
	UseMiddleware(Middleware) Application
	Handler() http.Handler
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Middleware func(http.Handler) http.Handler

type application struct {
	middleware []Middleware
	router     http.Handler
}

func NewApplication() Application {
	return new(application)
}

func (app *application) UseRouter(router http.Handler) Application {
	app.router = router
	return app
}

func (app *application) UseMiddleware(middleware Middleware) Application {
	if app.middleware == nil {
		app.middleware = make([]Middleware, 0)
	}

	app.middleware = append(app.middleware, middleware)

	return app
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := app.router

	if router == nil {
		router = noRouterHandler()
	}

	handler := reduce(app.middleware, router)

	handler.ServeHTTP(w, r)
}

func reduce(middleware []Middleware, acc http.Handler) http.Handler {
	if len(middleware) == 0 {
		return acc
	}

	return reduce(middleware[:len(middleware)-1], middleware[len(middleware)-1](acc))
}

func noRouter(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No router defined", http.StatusNotFound)
}

func noRouterHandler() http.Handler {
	return http.HandlerFunc(noRouter)
}
