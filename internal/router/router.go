package router

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

// 封装router
type Router struct {
	mux         *http.ServeMux
	middlewares []middleware
}

//NewRouter 创建一个新的Router实例，并且初始化Servemux和MIddleware切片

func NewRouter() *Router {
	return &Router{
		http.NewServeMux(),
		make([]middleware, 0),
	}
}

//添加中间件的处理函数

func (r *Router) Use(mw ...middleware) {
	r.middlewares = append(r.middlewares, mw...)
}

//路由注册函数

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

//实现handelr接口，在进入这个函数的时候现场进行中间件处理，然后调用handelr.SerHTTP来进入洋葱模型

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	handler.ServeHTTP(w, req)
}
