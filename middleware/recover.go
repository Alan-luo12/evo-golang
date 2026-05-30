package middlewares

import (
	"Lin/pkg/errors"
	"Lin/pkg/response"
	"log"
	"net/http"
)

// 捕获panic
func RecoverMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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
}
