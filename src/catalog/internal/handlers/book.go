package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"catalog/internal/service"
	"catalog/internal/shared"
	"catalog/internal/storage"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := storage.FetchBook(r.PathValue("id"), r.Context())
	if err != nil {
		shared.LogError(err)
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

func GetBookDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	book, err := storage.FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	if book == nil {
		w.WriteHeader(404)
		return
	}

	bytes, err := service.FetchBookDocument(book.FilePath)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res := service.EncodeDocument(bytes)
	w.Write([]byte(res))
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(400)
		return
	}
	defer r.Body.Close()
	var form service.BookPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(400)
		return
	}
}
