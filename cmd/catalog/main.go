package main

import (
	"log"
	"net/http"

	"github.com/madeinheaven91/patchouli/internal/catalog"
	"github.com/madeinheaven91/patchouli/internal/shared"
	rt "github.com/madeinheaven91/patchouli/pkg/router"
)

func main() {
	shared.InitFromEnv()
	catalog.InitDB()
	r := rt.NewRouter(http.NewServeMux())
	r.Route = rt.NewRoute("/").
		Subroute(
			rt.NewRoute("v1").
				Subroute(
					rt.NewRoute("/health").
						Method("GET").
						Handler(shared.HealthHandler{}),
				).
				Subroute(
					rt.NewRoute("/book/{hash}/info").
						Method("GET").
						Handler(catalog.GetBookInfo{}),
				),
		)
	log.Println("Starting server on port 3000")
	if err := http.ListenAndServe("0.0.0.0:3000", r.BuildMux()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
