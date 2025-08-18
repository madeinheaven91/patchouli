package main

import (
	"log"
	"net/http"

	hmw "github.com/madeinheaven91/hvnroutes/pkg/middlewares"
	r "github.com/madeinheaven91/hvnroutes/pkg/router"
	"github.com/madeinheaven91/patchouli/internal/catalog"
	"github.com/madeinheaven91/patchouli/internal/shared"
)

func main() {
	shared.InitFromEnv()
	catalog.InitDB()

	router := r.NewRouter(http.NewServeMux())
	router.Route = r.Route("/", nil, "").
		MiddlewareFunc(hmw.LoggingMW).
		Subroute(
			r.Route("v1").
				Subroute(
					r.Route("/book", nil, "POST").
						HandlerFunc(catalog.PostBook).
						MiddlewareFunc(shared.AuthMW).
						Subroute(
							r.Route("/{id}", nil, "GET", true).
								HandlerFunc(catalog.GetBook).
								MiddlewareFunc(hmw.LoggingMW).
								MiddlewareFunc(shared.AuthMW),
						).
						Subroute(
							r.Route("/{id}/document", nil, "GET", true).
								HandlerFunc(catalog.GetBookDocument).
								MiddlewareFunc(hmw.LoggingMW).
								MiddlewareFunc(shared.AuthMW),
						),
				).
				Subroute(
					r.Route("/category", nil, "POST").
						HandlerFunc(catalog.PostCategory).
						HandlerFunc(catalog.PostCategory).
						MiddlewareFunc(shared.AuthMW).
						Subroute(
							r.Route("/{id}", nil, "GET", true).
								HandlerFunc(catalog.GetCategory).
								MiddlewareFunc(hmw.LoggingMW).
								MiddlewareFunc(shared.AuthMW),
						),
				).
				Subroute(
					r.Route("/tag", nil, "POST").
						HandlerFunc(catalog.PostTag).
						MiddlewareFunc(shared.AuthMW).
						Subroute(
							r.Route("/{id}", nil, "GET", true).
								HandlerFunc(catalog.GetTag).
								MiddlewareFunc(hmw.LoggingMW).
								MiddlewareFunc(shared.AuthMW),
						),
				).
				Subroute(
					r.Route("/author", nil, "POST").
						HandlerFunc(catalog.PostAuthor).
						MiddlewareFunc(shared.AuthMW).
						Subroute(
							r.Route("/{id}", nil, "GET", true).
								HandlerFunc(catalog.GetAuthor).
								MiddlewareFunc(hmw.LoggingMW).
								MiddlewareFunc(shared.AuthMW),
						),
				),
		)

	log.Println("Starting server on port 3000")
	if err := http.ListenAndServe("0.0.0.0:3000", router.BuildMux()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
