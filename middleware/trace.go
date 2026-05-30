package middlewares

import (
	"Lin"
	"Lin/pkg/errors"
	"Lin/pkg/response"
	"Lin/pkg/snowid"
	"net/http"
	"strconv"
)

// 跟踪中间件
func TraceMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 从请求头中获取traceID
			traceID := r.Header.Get("X-Trace-ID")
			if traceID == "" {
				id, err := snowid.NextID()
				if err != nil {
					response.Error(w, errors.NewSystemError(50012, "Failed to generate trace ID", err))
					return
				}
				// 转换为字符串
				traceID = strconv.FormatUint(id, 10)
			}
			w.Header().Set("X-Trace-ID", traceID)
			// 将traceID添加到上下文
			ctx := Lin.WithTraceID(r.Context(), traceID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
