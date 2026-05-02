package main

import (
	"log"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriterWrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &ResponseWriterWrapper{
			w,
			http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)

		log.Printf(
			"[Success] method=%s path=%s status=%v latency=%s",
			r.Method,
			r.URL.Path,
			wrapper.status,
			time.Since(start),
		)
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println("Panic Recovered", err)
				Errorresponse(w, http.StatusInternalServerError, 0, "Internal Server Error")
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
