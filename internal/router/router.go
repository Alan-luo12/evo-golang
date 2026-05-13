package router

import (
	"net/http"
	"sort"
)

// middleware函数类型
type Middleware func(http.Handler) http.Handler

// 封装router
type Router struct {
	mux             *http.ServeMux
	middlewareitems []middlewareitem
}

const (
	RecoverPriority = 0
	TracePriority   = 10
	LogPriority     = 20
)

type middlewareitem struct {
	name     string
	mw       Middleware
	priority int
}

func NewRouter() *Router {
	return &Router{
		mux:             http.NewServeMux(),
		middlewareitems: make([]middlewareitem, 0),
	}
}

//添加中间件的处理函数

func (r *Router) Use(name string, mw Middleware, priority int) {
	r.middlewareitems = append(r.middlewareitems, middlewareitem{name,
		mw,
		priority})

	sort.Slice(r.middlewareitems, func(i, j int) bool {
		return r.middlewareitems[i].priority < r.middlewareitems[j].priority
	})
}

//路由注册函数

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

//实现handelr接口，在进入这个函数的时候现场进行中间件处理，然后调用handelr.SerHTTP来进入洋葱模型

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewareitems) - 1; i >= 0; i-- {
		handler = r.middlewareitems[i].mw(handler)
	}

	handler.ServeHTTP(w, req)
}

//11

// Chain 链式调用中间件
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// ChainFunc 链式调用中间件
func ChainFunc(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	ch := Chain(h, middlewares...)
	// 转换为.HandlerFunc类型,这里是类型断言,因为Chain返回的是http.Handler类型,而.HandlerFunc是http.Handler的实现类型
	return ch.ServeHTTP
}
