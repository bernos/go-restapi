package application

import "net/http"

type Application interface {
	Use(Middleware) Application
	ListenAndServe(string, http.Handler) error
	ListenAndServeTLS(string, string, string, http.Handler) error
}

type Middleware func(http.Handler) http.Handler

type application struct {
	middleware []Middleware
}

func NewApplication() Application {
	return new(application)
}

func (app *application) ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, reduce(app.middleware, handler))
}

func (app *application) ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

func (app *application) Use(middleware Middleware) Application {
	if app.middleware == nil {
		app.middleware = make([]Middleware, 0)
	}

	app.middleware = append(app.middleware, middleware)

	return app
}

func reduce(middleware []Middleware, acc http.Handler) http.Handler {
	if len(middleware) == 0 {
		return acc
	}

	return reduce(middleware[:len(middleware)-1], middleware[len(middleware)-1](acc))
}
