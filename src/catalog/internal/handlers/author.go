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
		shared.WriteError(w, 500, err)
		return
	}
	defer r.Body.Close()

	var form service.AuthorPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		shared.WriteError(w, 400, err)
		return
	}

	author, err := storage.AddAuthor(form, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	res, _ := json.Marshal(author)
	w.WriteHeader(201)
	w.Write(res)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := storage.FetchAllAuthors(r.Context())
	if err != nil {
		shared.WriteError(w, 500, err)
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
		shared.WriteError(w, 500, err)
		return
	}
	if author == nil {
		shared.WriteError(w, 404, err)
		return
	}

	res, _ := json.Marshal(author)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := storage.DeleteAuthor(id, r.Context())
	if err != nil {
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}
