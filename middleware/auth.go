package middlewares

import (
	"Lin"
	"Lin/pkg/errors"
	"Lin/pkg/response"
	"net/http"
	"strings"
)

func AuthMiddleware(requiretoken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 检查token是否匹配
			if requiretoken == "" {
				ctx := Lin.WithAuthSubject(r.Context(), "auth-disabled")
				// 换一个上下文，继续处理
				next.ServeHTTP(w, r.WithContext(ctx))
				// 处理完成后，返回，避免继续处理
				return
			}

			tk := r.Header.Get("X-AUTH-TOKEN")
			if tk == "" {
				authHeader := r.Header.Get("Authorization")
				const (
					prefix = "Bearer "
				)
				if strings.HasPrefix(authHeader, prefix) {
					tk = strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
				}
			}

			if tk != requiretoken {
				response.Error(w, errors.NewUnauthorizedError(40100, "Unauthorized", nil))
				return
			}

			ctx := Lin.WithAuthSubject(r.Context(), "authorized")
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
