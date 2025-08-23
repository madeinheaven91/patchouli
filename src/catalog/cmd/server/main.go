package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"catalog/internal/handlers"
	"catalog/internal/service"
	"catalog/internal/shared"
	"catalog/internal/storage"

	"github.com/joho/godotenv"
	"github.com/madeinheaven91/hvnroutes/pkg/middlewares"
	r "github.com/madeinheaven91/hvnroutes/pkg/router"
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
				Route("/{name}", r.WrapHandler(handlers.FetchBookFile), "GET", logging, true).
				Route("/{name}", r.WrapHandler(handlers.DeleteBookFile), "DELETE").
				Route("/{name}/info", r.WrapHandler(handlers.FetchBookFileInfo), "GET"),
		).
		Service(
			r.NewRoute("/book", r.WrapHandler(handlers.GetAllBooks), "GET").
				Route("/{id}", r.WrapHandler(handlers.GetBook), "GET").
				Route("/{id}", r.WrapHandler(handlers.DeleteBook), "DELETE", auth, false).
				Route("/{id}/tags", r.WrapHandler(handlers.GetBookTags), "GET").
				Route("/{id}/tags/{tag}", r.WrapHandler(handlers.PostBookTag), "POST", auth, false).
				Route("/{id}/tags/{tag}", r.WrapHandler(handlers.DeleteBookTag), "DELETE", auth, false),
		).
		Route("/request", r.WrapHandler(handlers.GetAllRequests), "GET", auth, false).
		Service(
			r.NewRoute("/request", r.WrapHandler(handlers.PostRequest), "POST", auth, false).
				Route("/{id}", r.WrapHandler(handlers.GetRequest), "GET").
				Route("/{id}", r.WrapHandler(handlers.DeleteRequest), "DELETE").
				Route("/{id}/tags", r.WrapHandler(handlers.GetRequestTags), "GET").
				Route("/{id}/tags/{tag}", r.WrapHandler(handlers.PostRequestTag), "POST").
				Route("/{id}/tags/{tag}", r.WrapHandler(handlers.DeleteRequestTag), "DELETE").
				Route("/{id}/publish", r.WrapHandler(handlers.PublishRequest), "POST"),
		).
		Service(
			r.NewRoute("/tag", r.WrapHandler(handlers.PostTag), "POST", auth, false).
				Route("", r.WrapHandler(handlers.GetAllTags), "GET", logging, true).
				Route("/{name}", r.WrapHandler(handlers.DeleteTag), "DELETE").
				Route("/{name}", r.WrapHandler(handlers.GetTag), "GET", logging, true),
		).
		Service(
			r.NewRoute("/author", r.WrapHandler(handlers.PostAuthor), "POST", auth, false).
				Route("", r.WrapHandler(handlers.GetAllAuthors), "GET", logging, true).
				Route("/{id}", r.WrapHandler(handlers.DeleteAuthor), "DELETE").
				Route("/{id}", r.WrapHandler(handlers.GetAuthor), "GET", logging, true),
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
