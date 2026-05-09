package router

import (
	"log"
	"net/http"
	"time"

	"app/pkg/errors"
	"app/pkg/response"
)

//ResponseWriterWrapper结构体
type ResponseWriterWrapper struct {
	http.ResponseWriter
	status int
}

//写入响应头
func (r *ResponseWriterWrapper) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
}

//捕获日志

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &ResponseWriterWrapper{w, 200}

		next.ServeHTTP(wrapper, r)

		log.Printf(
			"%s %s status=%d latency=%s",
			r.Method,
			r.URL.Path,
			wrapper.status,
			time.Since(start),
		)
	})
}

//捕获panic
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			err := recover()
			if err != nil {
				log.Println("[Panic Recovered]", err)
				response.Error(w, errors.NewSystemError(5002, "Manual Panic triggered", nil))
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
