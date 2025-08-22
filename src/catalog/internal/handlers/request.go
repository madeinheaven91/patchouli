package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"catalog/internal/service"
	"catalog/internal/shared"
	"catalog/internal/storage"
)

func PostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	var form service.RequestPostForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		shared.LogError(err, string(body))
		w.WriteHeader(400)
		return
	}

	// form filename is a temporary generated name
	// replace it with book author + title
	oldFilename := form.Filename
	suffix := strings.Split(oldFilename, ".")[1]
	newFilename := shared.ToFilename(form.AuthorName+"_"+form.Title) + "." + suffix
	_, err = service.RenameBook(newFilename, oldFilename, r)
	if err != nil {
		shared.LogError(err, fmt.Sprintf("%#v", form), string(body), newFilename)
		w.WriteHeader(500)
		return
	}
	form.Filename = newFilename

	request, err := storage.AddRequest(form, r.Context())
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	for _, tag := range form.Tags {
		storage.AddTagToRequest(tag, request.ID, r.Context())
	}

	// might be incorrect
	req := service.RequestResponseFromModel(*request, form.Tags)
	res, _ := json.Marshal(req)
	w.WriteHeader(201)
	w.Write(res)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	err = storage.Delete("request", "id", id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	err = service.DeleteBook(req.Filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	tags, err := storage.FetchRequestTags(req.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	res := service.RequestResponseFromModel(*req, tags)
	json, _ := json.Marshal(res)
	w.Write(json)
}

func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	reqs, err := storage.FetchAllRequests(r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	res := make([]service.RequestResponse, 0)
	for _, req := range reqs {
		tags, err := storage.FetchRequestTags(req.ID, r.Context())
		if err != nil {
			shared.LogError(err)
			shared.WriteError(w, 500, err)
			return
		}
		resp := service.RequestResponseFromModel(req, tags)
		res = append(res, resp)
	}

	json, _ := json.Marshal(res)
	w.Write(json)
}

func PublishRequest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	tags, err := storage.FetchRequestTags(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	book, err := storage.AddBook(*req, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	for _, tag := range tags {
		storage.AddTagToBook(tag.Name, book.ID, r.Context())
	}

	err = storage.Delete("request", "id", id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(201)
}

func GetRequestTags(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	tags, err := storage.FetchRequestTags(req.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	json, _ := json.Marshal(tags)
	w.Write(json)
}

func PostRequestTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tag := r.URL.Query().Get("tag")

	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	_, err = storage.AddTagToRequest(tag, req.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(201)
}

func DeleteRequestTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tag := r.URL.Query().Get("tag")

	req, err := storage.FetchRequest(id, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	err = storage.DeleteTagToRequest(tag, req.ID, r.Context())
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	w.WriteHeader(204)
}
