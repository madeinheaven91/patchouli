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
	path      string
	handler   http.Handler
	method    string
	subroutes []*Route
}

type defaultHandler struct{}

func (d defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func NewRoute(path string) *Route {
	return &Route{
		path,
		defaultHandler{},
		"",
		[]*Route{},
	}
}

func (r Route) buildSubroutes(mux *http.ServeMux, prefix string) {
	newPath := prefix + r.path
	mux.Handle(r.method+" "+newPath, r.handler)
	for _, route := range r.subroutes {
		go route.buildSubroutes(mux, newPath)
	}
}

func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

func (r *Route) Handler(handler http.Handler) *Route {
	r.handler = handler
	return r
}

func (r *Route) Method(method string) *Route {
	r.method = method
	return r
}

func (r *Route) Subroute(subroute *Route) *Route {
	r.subroutes = append(r.subroutes, subroute)
	return r
}
