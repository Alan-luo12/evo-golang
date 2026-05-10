package router

import (
	"log"
	"myapp/internal/pkg"
	"net/http"
	"strconv"
	"time"

	"myapp/pkg/errors"
	"myapp/pkg/response"
	"myapp/pkg/snowid"
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
			pkg.GetTraceID(r.Context()),
		)
	})
}

// 捕获panic
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			err := recover()

			if err != nil {

				id := w.Header().Get("X-Trace-ID")
				if id == "" {
					id = "unknown"
				}
				log.Println("[Panic Recovered]", id, "err:", err)
				response.Error(w, errors.NewSystemError(5004, "Internal Server Error", nil))
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取traceID
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			id, _ := snowid.NextID()
			// 转换为字符串
			traceID = strconv.FormatUint(id, 10)
		}
		// 将traceID添加到上下文
		ctx := pkg.WithTraceID(r.Context(), traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
