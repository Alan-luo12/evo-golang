package router

import (
	"app/internal/limit"
	"app/pkg/errors"
	"app/pkg/response"
	"log"
	"net/http"
)

//这里做了特别改动，该成立可以接受HandlerFunc类型和返回HandlerFunc类型的函数，专门用来适配包装原来的普通函数，做一个限流中间件，因为普通
//函数能够被http.HandlerFunc转换类型未http.HandlerFunc类型，而http.HandlerFunc类型又实现了http.Handler接口，
//所以可以适配原来的HandleFunc函数。

func NewRateLimitMiddleware(tb *limit.TokenBucket) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !tb.Allow() {
				response.Error(w, errors.NewLimitExceededError(4007, "Rate Limit Exceeded", nil))
				log.Println("[Limit Excedeed] Rate Limit Exceeded")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
