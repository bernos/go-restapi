package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func LoggerMiddleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	start := time.Now()

	next.ServeHTTP(w, r)

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}
