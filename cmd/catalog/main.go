package main

import (
	"log"
	"net/http"

	"github.com/madeinheaven91/patchouli/internal/catalog"
	"github.com/madeinheaven91/patchouli/internal/shared"
	r "github.com/madeinheaven91/patchouli/pkg/router"
)

func main() {
	shared.InitFromEnv()
	catalog.InitDB()

	router := r.NewRouter(http.NewServeMux())
	router.Route = r.NewRoute("/v1").
		MiddlewareFunc(shared.LoggingMW).
		Subroute(
			r.NewRoute("/book").
				Method("POST").
				MiddlewareFunc(shared.AuthMW).
				HandlerFunc(catalog.PostBook).
				Subroute(
					r.NewRoute("/{id}").
						Method("GET").
						StopMiddleware().
						MiddlewareFunc(shared.LoggingMW).
						HandlerFunc(catalog.GetBook),
				).
				Subroute(
					r.NewRoute("/{id}/document").
						Method("GET").
						StopMiddleware().
						MiddlewareFunc(shared.LoggingMW).
						HandlerFunc(catalog.GetBookDocument),
				),
		).
		Subroute(
			r.NewRoute("/category").
				Method("POST").
				MiddlewareFunc(shared.AuthMW).
				HandlerFunc(catalog.PostCategory).
				Subroute(
					r.NewRoute("/{id}").
						Method("GET").
						StopMiddleware().
						MiddlewareFunc(shared.LoggingMW).
						MiddlewareFunc(shared.AuthMW).
						HandlerFunc(catalog.GetCategory),
				),
		).
		Subroute(
			r.NewRoute("/tag").
				Method("POST").
				MiddlewareFunc(shared.AuthMW).
				HandlerFunc(catalog.PostTag).
				Subroute(
					r.NewRoute("/{id}").
						Method("GET").
						StopMiddleware().
						MiddlewareFunc(shared.LoggingMW).
						MiddlewareFunc(shared.AuthMW).
						HandlerFunc(catalog.GetTag),
				),
		).
		Subroute(
			r.NewRoute("/author").
				Method("POST").
				MiddlewareFunc(shared.AuthMW).
				HandlerFunc(catalog.PostAuthor).
				Subroute(
					r.NewRoute("/{id}").
						Method("GET").
						StopMiddleware().
						MiddlewareFunc(shared.LoggingMW).
						MiddlewareFunc(shared.AuthMW).
						HandlerFunc(catalog.GetAuthor),
				),
		)

	log.Println("Starting server on port 3000")
	if err := http.ListenAndServe("0.0.0.0:3000", router.BuildMux()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
