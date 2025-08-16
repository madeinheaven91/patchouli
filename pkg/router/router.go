// Package router
//
// Self made routing because frameworks are gay
package router

import (
	"net/http"
)

type Router struct {
	Mux   *http.ServeMux
	Route *Route
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{
		Mux: mux,
	}
}

func (c *Router) BuildMux() *http.ServeMux {
	c.Route.buildSubroutes(c.Mux, "")
	return c.Mux
}

type Route struct {
	path            string
	handler         http.HandlerFunc
	method          string
	subroutes       []*Route
	middlewareChain []Middleware
	stopMiddleware  bool
}

type defaultHandler struct{}

func (d defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func NewRoute(path string) *Route {
	return &Route{
		path,
		defaultHandler{}.ServeHTTP,
		"",
		[]*Route{},
		nil,
		false,
	}
}

func (rt *Route) buildSubroutes(mux *http.ServeMux, prefix string) {
	newPath := prefix + rt.path
	mux.Handle(rt.method+" "+newPath, rt.buildChain())
	// spew.Dump(rt.method+" "+newPath)
	// spew.Dump(rt)

	for _, route := range rt.subroutes {
		if !route.stopMiddleware || rt.middlewareChain == nil {
			route.middlewareChain = append(rt.middlewareChain, route.middlewareChain...)
		}
		// TODO: implement middleware propagation
		go route.buildSubroutes(mux, newPath)
	}
}

func (rt *Route) Path(path string) *Route {
	rt.path = path
	return rt
}

func (rt *Route) Handler(handler http.Handler) *Route {
	rt.handler = handler.ServeHTTP
	return rt
}

func (rt *Route) HandlerFunc(handler http.HandlerFunc) *Route {
	rt.handler = handler
	return rt
}

func (rt *Route) Method(method string) *Route {
	rt.method = method
	return rt
}

func (rt *Route) Subroute(subroute *Route) *Route {
	rt.subroutes = append(rt.subroutes, subroute)
	return rt
}

func (rt *Route) Middleware(mw Middleware) *Route {
	rt.middlewareChain = append(rt.middlewareChain, mw)
	return rt
}

func (rt *Route) MiddlewareFunc(mw MiddlewareFunc) *Route {
	rt.middlewareChain = append(rt.middlewareChain, basicWrapper{mw})
	return rt
}

func (rt *Route) StopMiddleware() *Route {
	rt.stopMiddleware = true
	return rt
}
