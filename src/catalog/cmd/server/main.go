package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/madeinheaven91/hvnroutes/pkg/middlewares"
	r "github.com/madeinheaven91/hvnroutes/pkg/router"
	"github.com/madeinheaven91/patchouli/internal/handlers"
	"github.com/madeinheaven91/patchouli/internal/service"
	"github.com/madeinheaven91/patchouli/internal/shared"
	"github.com/madeinheaven91/patchouli/internal/storage"
)

func main() {
	er := godotenv.Load(".env")
	if er != nil {
		shared.LogError(er)
		os.Exit(1)
	}
	shared.InitFromEnv()
	storage.Init()
	err := service.InitMinio()
	if err != nil {
		shared.LogError(err)
		os.Exit(1)
	}

	auth := r.WrapMW(shared.AuthMW)
	logging := middlewares.Logging{}

	router := r.NewRouter(http.NewServeMux())
	router.Route = r.NewRoute("/v1", nil, "", logging, false).
		Service(
			r.NewRoute("/files", r.WrapHandler(handlers.UploadBookFile), "POST", auth, false).
				Route("/{name}", r.WrapHandler(handlers.FetchBookFile), "GET").
				Route("/{name}/info", r.WrapHandler(handlers.FetchBookFileInfo), "GET"),
		).
		Service(
			r.NewRoute("/book", r.WrapHandler(handlers.PostBook), "POST", auth, false).
				Route("/{id}", nil, "GET", logging, true).
				Route("/{id}/document", nil, "GET", logging, true),
		).
		Service(
			r.NewRoute("/request", r.WrapHandler(handlers.PostRequest), "POST", auth, false).
				Route("/{id}", nil, "GET", logging, true).
				Route("/{id}", nil, "DELETE").
				// Route("/{id}/tags", r.WrapHandler(handlers.PostTag), "POST").
				Route("/{id}/publish", nil, "POST"),
		).
		Service(
			r.NewRoute("/tag", r.WrapHandler(handlers.PostTag), "POST", auth, false).
				Route("/{name}", r.WrapHandler(handlers.DeleteTag), "DELETE").
				Route("/{name}", r.WrapHandler(handlers.GetTag), "GET", logging, true).
				Route("/all", r.WrapHandler(handlers.GetAllTags), "GET", logging, true),
		).
		Service(
			r.NewRoute("/author", r.WrapHandler(handlers.PostAuthor), "POST", auth, false).
				Route("/{id}", r.WrapHandler(handlers.DeleteAuthor), "DELETE").
				Route("/{id}", r.WrapHandler(handlers.GetAuthor), "GET", logging, true).
				Route("/all", r.WrapHandler(handlers.GetAllAuthors), "GET", logging, true),
		)

	server := http.Server{
		Addr:        ":3000",
		IdleTimeout: 120 * time.Second,
		Handler:     router.BuildMux(),
	}

	go func() {
		log.Println("Starting server on port 3000")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	log.Printf("%v, shutting down\n", sig)

	ctx := context.Background()
	server.Shutdown(ctx)
	storage.Close()
}
