package Lin

import "net/http"

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
