package middlewares

import (
	"Lin"
	"log"
	"net/http"
	"time"
)

// ResponseWriterWrapper结构体
type ResponseWriterWrapper struct {
	http.ResponseWriter
	status int
}

// 写入响应头
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
			"%s %s status=%d latency=%s traceId:%s",
			r.Method,
			r.URL.Path,
			wrapper.status,
			time.Since(start),
			Lin.GetTraceID(r.Context()),
		)
	})
}
