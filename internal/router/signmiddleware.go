package router

import (
	"bytes"
	"io"
	"myapp/internal/app/repo"
	"myapp/internal/pkg"
	"myapp/pkg/errors"
	"myapp/pkg/response"
	"net/http"
	"strconv"
	"time"
)

func NewSignMiddleware(secret string, window time.Duration, redisrepo *repo.RedisRepo) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sign := r.Header.Get("X-Sign")
			tsstr := r.Header.Get("X-TimeStamp")
			nonce := r.Header.Get("X-Nonce")

			// 验证Sign、TimeStamp、Nonce是否为空,获取ts还有ttl
			if sign == "" || tsstr == "" || nonce == "" {
				response.Error(w, errors.NewConflictError(4091, "Sign, TimeStamp, or Nonce is empty", nil))
				return
			}
			ts, err := strconv.ParseInt(tsstr, 10, 64)
			if err != nil {
				response.Error(w, errors.NewConflictError(4092, "TimeStamp is invalid", nil))
				return
			}
			ttl := pkg.TimeStampTTL(time.Now(), ts, window)

			// 验证Nonce是否有效
			if !pkg.ValidNonce(nonce) {
				response.Error(w, errors.NewConflictError(4093, "Nonce is invalid", nil))
				return
			}
			// 验证TimeStamp是否在窗口内
			if !pkg.TimeStampInWindow(time.Now(), ts, window) {
				response.Error(w, errors.NewConflictError(4094, "TimeStamp is not in window", nil))
				return
			}

			// 验证Sign是否有效
			body, err := io.ReadAll(r.Body)
			if err != nil {
				response.Error(w, errors.NewConflictError(4095, "Read body body failed", err))
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(body))
			payload := pkg.BuildSignPayload(r.Method, r.URL.Path, ts, nonce, body)
			if !pkg.VertifySign(payload, secret, sign) {
				response.Error(w, errors.NewConflictError(4096, "Sign is invalid", nil))
				return
			}

			// 验证Nonce是否被使用过
			success, err := redisrepo.UseNonceOnce(r.Context(), nonce, ttl)
			if err != nil {
				response.Error(w, errors.NewConflictError(4097, "UseNonceOnce failed", err))
				return
			}
			if !success {
				response.Error(w, errors.NewConflictError(4098, "Nonce is invalid", nil))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
