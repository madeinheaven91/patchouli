// Package catalog
//
// Defines all the stuff for the book catalog microservice
package catalog

import (
	"encoding/json"
	"log"
	"net/http"
)

// Endpoints:
//
// get book
// post book
// patch book
// delete book

type GetBookInfo struct{}

func (g GetBookInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler fired")
	hash := r.PathValue("hash")
	book, err := FetchBook(hash, r.Context())
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(500)
		return
	}
	if book == nil {
		w.WriteHeader(404)
		return
	}
	res, _ := json.Marshal(book)
	w.Write(res)
}
