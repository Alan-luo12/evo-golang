package router

import (
	"context"
	"fmt"
	"myapp/internal/app/repo"
	"myapp/pkg/errors"
	"myapp/pkg/response"
	"net/http"
	"time"
)

// NewDistLimitMiddleware 新的dist limit中间件
func NewDistLimitMiddleware(ctx context.Context, repo *repo.RedisRepo, max int64, window time.Duration, failopen bool) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Path
			//如果max小于等于0，直接放行
			if max <= 0 {
				next.ServeHTTP(w, r)
				return
			}

			//如果window小于等于0，直接放行
			if window <= 0 {
				window = time.Second
			}

			bucketid := time.Now().UnixMilli() / window.Milliseconds()
			key := fmt.Sprintf("ratelimit:%s:%d", name, bucketid)

			//检查是否超过限制
			res, err := repo.AllowDist(ctx, key, window, max)

			if err != nil {
				if failopen {
					//在这里不需要设置关于limit的响应头，因为failopen是true，所以直接返回错误信息
					w.Header().Set("X-Dist-Limit-Mode", "redis-failopen")
					next.ServeHTTP(w, r)
					return
				}
				response.Error(w, errors.NewSystemError(50010, "Failed to check the dist limit", err))
				return
			}

			remain := max - res.Current

			//设置响应头在判断是否超过限制之前
			w.Header().Set("X-Dist-Limit-Mode", "redis")
			w.Header().Set("X-Dist-Limit-Remain", fmt.Sprintf("%d", remain))
			w.Header().Set("X-Dist-Limit-Window", fmt.Sprintf("%d", window.Milliseconds()))
			w.Header().Set("X-Dist-Limit-TTL", fmt.Sprintf("%d", res.TTL.Milliseconds()))

			if !res.Allow {
				response.Error(w, errors.NewLimitExceededError(50011, "Dist Limit Exceeded", nil))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
