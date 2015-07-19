package middleware

import "net/http"

type MiddlewareStack struct {
	middleware []Middleware
	handler    http.Handler
}

type Middleware func(http.Handler) http.Handler

func New() *MiddlewareStack {
	return &MiddlewareStack{
		middleware: make([]Middleware, 0),
	}
}

func (stack *MiddlewareStack) Use(middleware Middleware) *MiddlewareStack {
	stack.middleware = append(stack.middleware, middleware)

	return stack
}

func (stack *MiddlewareStack) UseHandler(h http.Handler) *MiddlewareStack {
	stack.handler = h
	return stack
}

func MiddlewareFunc(m func(http.ResponseWriter, *http.Request, http.Handler)) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m(w, r, h)
		})
	}
}

func (stack *MiddlewareStack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := stack.handler

	if handler == nil {
		handler = http.HandlerFunc(defaultHandler)
	}

	pipeline := reduce(stack.middleware, handler)

	pipeline.ServeHTTP(w, r)
}

func reduce(middleware []Middleware, acc http.Handler) http.Handler {
	if len(middleware) == 0 {
		return acc
	}

	return reduce(middleware[:len(middleware)-1], middleware[len(middleware)-1](acc))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

}
