package main

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler
type Router struct {
	mux         *http.ServeMux
	middlewares []middleware
}

func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		middlewares: make([]middleware, 0),
	}
}

func (r *Router) Use(mw ...middleware) {
	r.middlewares = append(r.middlewares, mw...)
}

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	handler.ServeHTTP(w, req)
}
