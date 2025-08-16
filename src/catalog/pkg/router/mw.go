package router

import (
	"net/http"
	"slices"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

type Middleware interface {
	Handle(http.HandlerFunc) http.HandlerFunc
}

type basicWrapper struct {
	fn MiddlewareFunc
}

func (b basicWrapper) Handle(next http.HandlerFunc) http.HandlerFunc {
	return b.fn(next)
}

func (rt *Route) buildChain() http.HandlerFunc {
	chain := func(handler http.HandlerFunc) http.HandlerFunc {
		for _, mw := range slices.Backward(rt.middlewareChain) {
			handler = mw.Handle(handler)
		}
		return handler
	}
	return chain(rt.handler)
}
