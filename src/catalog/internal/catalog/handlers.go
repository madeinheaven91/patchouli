// Package catalog
//
// Defines all the stuff for the book catalog microservice
package catalog

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/madeinheaven91/patchouli/internal/shared"
)

// Endpoints:
//
// get book
// post book
// delete book

func GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := FetchBook(r.PathValue("id"), r.Context())
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
	book, err := FetchBook(id, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	if book == nil {
		w.WriteHeader(404)
		return
	}

	bytes, err := FetchBookDocument(book.FilePath)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res := EncodeDocument(bytes)
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
	var form BookPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(400)
		return
	}
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	var form struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		w.WriteHeader(400)
		return
	}

	category, err := AddCategory(form.Name, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(category)
	w.WriteHeader(201)
	w.Write(res)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(404)
		return
	}
	tag, err := FetchTag(id, r.Context())
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if tag == nil {
		w.WriteHeader(404)
		return
	}

	res, _ := json.Marshal(tag)
	w.Write(res)
}

func PostTag(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	var form struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		w.WriteHeader(400)
		return
	}

	tag, err := AddTag(form.Name, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(tag)
	w.WriteHeader(201)
	w.Write(res)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(404)
		return
	}
	category, err := FetchCategory(id, r.Context())
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if category == nil {
		w.WriteHeader(404)
		return
	}

	res, _ := json.Marshal(category)
	w.Write(res)
}

func PostAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	var form AuthorPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		w.WriteHeader(400)
		return
	}

	author, err := AddAuthor(form, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(author)
	w.WriteHeader(201)
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(404)
		return
	}
	author, err := FetchAuthor(id, r.Context())
	if err != nil {
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
