v10 进行重构 从“异步任务服务”变成了一个真正可复用的轻量 Web 框架
快速开始
GO
package main
import (
	"Lin"
	"Lin/middlewares"
	"Lin/pkg/response"
	"Lin/pkg/snowid"
	"Lin/ratelimit"
	"net/http"
)
func main() {
	// TraceMiddleware 依赖 snowid
	snowid.Init(1)
	app := Lin.NewRouter()
	tb := ratelimit.NewTokenBucket(100, 10)
	app.Use("recover", middlewares.RecoverMiddleware(), Lin.RecoverPriority)
	app.Use("trace", middlewares.TraceMiddleware(), Lin.TracePriority)
	app.Use("log", middlewares.LogMiddleware, Lin.LogPriority)
	app.Use("auth", middlewares.AuthMiddleware("your-token"), 30)
	app.Use("limit", middlewares.LocalLimitMiddleware(tb), 40)
	app.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, map[string]any{
			"message": "pong",
			"traceId": Lin.GetTraceID(r.Context()),
			"auth":    Lin.GetAuthSubject(r.Context()),
		})
	})
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
统一响应
成功响应：

JSON
{
  "code": 0,
  "msg": "Success",
  "data": {}
}
错误响应：

JSON
{
  "code": 40100,
  "msg": "Unauthorized",
  "data": null
}
错误处理
GO
response.Error(w, errors.NewUserError(4001, "invalid params", err))
response.Error(w, errors.NewSystemError(5001, "server error", err))
response.Error(w, errors.NewUnauthorizedError(40100, "Unauthorized", nil))
response.Error(w, errors.NewLimitExceededError(4007, "Rate Limit Exceeded", nil))
错误类型会自动映射 HTTP 状态码：

ErrorType	HTTP Status
User	400
System	500
LimitExceeded	429
NotFound	404
Unauthorized	401
Conflict	409
中间件
Recover
捕获 panic，返回统一 500 响应。

GO
app.Use("recover", middlewares.RecoverMiddleware(), Lin.RecoverPriority)
Trace
自动生成或透传 X-Trace-ID。

GO
snowid.Init(1)
app.Use("trace", middlewares.TraceMiddleware(), Lin.TracePriority)
Log
记录请求方法、路径、状态码、耗时和 TraceID。

GO
app.Use("log", middlewares.LogMiddleware, Lin.LogPriority)
Auth
支持：

HTTP
X-AUTH-TOKEN: your-token
或：

HTTP
Authorization: Bearer your-token
使用：

GO
app.Use("auth", middlewares.AuthMiddleware("your-token"), 30)
如果 token 为空，则跳过鉴权：

GO
middlewares.AuthMiddleware("")
LocalLimit
本地令牌桶限流：

GO
tb := ratelimit.NewTokenBucket(100, 10)
app.Use("limit", middlewares.LocalLimitMiddleware(tb), 40)
Sign
基于 HMAC-SHA256 的签名校验，防止请求伪造和重放。

请求头：

HTTP
X-Sign: <signature>
X-TimeStamp: <unix timestamp>
X-Nonce: <random nonce>
使用：

GO
store := security.NewMemoryNonceStore()
app.Use("sign", middlewares.SignMiddleware(
	"your-secret",
	time.Minute*5,
	store,
), 50)
签名 payload 由以下内容组成：

TEXT
method + path + timestamp + nonce + body
Context
获取 TraceID：

GO
traceID := Lin.GetTraceID(r.Context())
获取鉴权主体：

GO
subject := Lin.GetAuthSubject(r.Context())
Queue / WorkerPool
GO
package main
import (
	"Lin/queue"
	"Lin/works"
	"context"
	"log"
)
func main() {
	ctx := context.Background()
	q := queue.NewMemoryQueue(100)
	dispatcher := works.NewDispatcher(q, 100)
	dispatcher.Start(ctx)
	pool := works.NewWorkerPool(3, dispatcher.Jobs(), works.ProcessFunc(
		func(ctx context.Context, job queue.Job) error {
			log.Println("process job:", job.ID, job.Name)
			return nil
		},
	))
	pool.Start(ctx)
	_ = q.Enqueue(ctx, queue.Job{
		ID:      "1",
		Name:    "demo",
		Payload: []byte("hello"),
	})
	pool.Wait()
}
中间件优先级
优先级数字越小，越靠外层执行。

默认：

GO
const (
	RecoverPriority = 0
	TracePriority   = 10
	LogPriority     = 20
)
推荐顺序：

TEXT
Recover -> Trace -> Log -> Auth -> Limit -> Sign -> Handler
说明
Lin 适合作为学习型或小型项目基础框架使用，核心设计保持简单，主要依赖 Go 标准库 net/http。