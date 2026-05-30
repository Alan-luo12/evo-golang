package Lin

import (
	"context"
	"net/http"
	"sort"
	"time"
)

// middleware函数类型
type Middleware func(http.Handler) http.Handler

// router加中间件管理加http.Server
type App struct {
	mux             *http.ServeMux
	middlewareitems []middlewareitem
	server          *http.Server
}

// 中间件优先级
const (
	RecoverPriority = 0
	TracePriority   = 10
	LogPriority     = 20
)

// 中间件项
type middlewareitem struct {
	name     string
	mw       Middleware
	priority int
}

// 创建路由
func NewRouter() *App {
	return &App{
		mux:             http.NewServeMux(),
		middlewareitems: make([]middlewareitem, 0),
		server:          nil,
	}
}

// 添加中间件的处理函数
func (r *App) Use(name string, mw Middleware, priority int) {
	r.middlewareitems = append(r.middlewareitems, middlewareitem{name,
		mw,
		priority})

	sort.Slice(r.middlewareitems, func(i, j int) bool {
		return r.middlewareitems[i].priority < r.middlewareitems[j].priority
	})
}

//路由注册函数

func (r *App) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

//实现handelr接口，在进入这个函数的时候现场进行中间件处理，然后调用handelr.SerHTTP来进入洋葱模型

func (r *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewareitems) - 1; i >= 0; i-- {
		handler = r.middlewareitems[i].mw(handler)
	}

	handler.ServeHTTP(w, req)
}

// 启动路由
func (r *App) Run(addr string) error {
	if addr == "" {
		addr = ":8080"
	}
	r.server = &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}
	return r.server.ListenAndServe()
}

// 停止路由
func (r *App) Stop(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
