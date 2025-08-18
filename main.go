package main

import (
	"log"
	"net/http"

	mw "github.com/madeinheaven91/hvnroutes/pkg/middlewares"
	r "github.com/madeinheaven91/hvnroutes/pkg/router"
)

func main() {
	router := r.NewRouter(http.NewServeMux())
	router.Route = r.Route("/", nil, "GET").
		MiddlewareFunc(mw.LoggingMW).
		HandlerFunc(hdl)

	if err := http.ListenAndServe(":9000", router.BuildMux()); err != nil {
		log.Println(err)
	}
}

func hdl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(666)
}
