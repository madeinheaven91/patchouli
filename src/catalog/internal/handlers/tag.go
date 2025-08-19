package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"catalog/internal/models"
	"catalog/internal/shared"
	"catalog/internal/storage"
)

func GetTag(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	tag, err := storage.FetchTag(name, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	if tag == nil {
		shared.WriteError(w, 404, "tag not found")
		return
	}

	res, _ := json.Marshal(tag)
	w.Write(res)
}

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	tag, err := storage.FetchAllTags(r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	res, _ := json.Marshal(tag)
	w.Write(res)
}

func PostTag(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	defer r.Body.Close()

	var form models.Tag
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 400, "invalid json")
		return
	}

	tag, err := storage.AddTag(form.Name, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	res, _ := json.Marshal(tag)
	w.WriteHeader(201)
	w.Write(res)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	err := storage.Delete("tag", "name", name, r.Context())
	// err := storage.DeleteTag(name, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}
