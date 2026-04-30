package main

import (
	"log"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	Status int
}

func (w *ResponseWriterWrapper) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}

//日志写入中间件

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &ResponseWriterWrapper{
			w,
			200,
		}

		next.ServeHTTP(wrapper, r)

		log.Printf(
			"%s %s status=%d time=%s",
			r.Method,
			r.URL.Path,
			wrapper.Status,
			time.Since(start),
		)
	})
}

//防崩溃中间件

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Panic recovered", err)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
