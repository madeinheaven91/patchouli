package handlers

import (
	"encoding/json"
	"net/http"

	"catalog/internal/service"
	"catalog/internal/shared"
	"catalog/internal/storage"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := storage.FetchBook(r.PathValue("id"), r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	// This shouldn't work, but anyway
	if book == nil {
		shared.WriteError(w, 404, err)
		return
	}

	tags, err := storage.FetchBookTags(book.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	res := service.BookResponseFromModel(*book, tags)
	json, _ := json.Marshal(res)
	w.Write(json)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := storage.FetchAllBooks(r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	res := make([]service.BookResponse, 0)
	for _, book := range books {
		tags, err := storage.FetchBookTags(book.ID, r.Context())
		if err != nil {
			shared.LogError(err)
			shared.WriteError(w, 500, err)
			return
		}
		resp := service.BookResponseFromModel(book, tags)
		res = append(res, resp)
	}

	json, _ := json.Marshal(res)
	w.Write(json)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	book, err := storage.FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	err = storage.Delete("book", "id", id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	err = service.DeleteBook(book.Filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}

func GetBookTags(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	book, err := storage.FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	tags, err := storage.FetchBookTags(book.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	json, _ := json.Marshal(tags)
	w.Write(json)
}

func PostBookTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tag := r.PathValue("tag")

	book, err := storage.FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	_, err = storage.AddTagToBook(tag, book.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(201)
}

func DeleteBookTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tag := r.PathValue("tag")

	book, err := storage.FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	err = storage.DeleteTagToBook(tag, book.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}
