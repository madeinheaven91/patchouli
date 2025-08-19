package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"catalog/internal/service"
	"catalog/internal/shared"
	"catalog/internal/storage"
)

func PostAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	var form service.AuthorPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		w.WriteHeader(400)
		return
	}

	author, err := storage.AddAuthor(form, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(author)
	w.WriteHeader(201)
	w.Write(res)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := storage.FetchAllAuthors(r.Context())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	res, _ := json.Marshal(authors)
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	author, err := storage.FetchAuthor(id, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	if author == nil {
		w.WriteHeader(404)
		return
	}

	res, _ := json.Marshal(author)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := storage.DeleteAuthor(id, r.Context())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
