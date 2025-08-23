package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"catalog/internal/service"
	"catalog/internal/shared"
)

// NOTE: im not sure whether multipart upload is required, maybe ill implement it later
// for now i will use a simple solution. Maybe its bad, idk

// UploadBookFile handles a file upload in single http request.
func UploadBookFile(w http.ResponseWriter, r *http.Request) {
	filename, err := service.UploadBook(r)
	if err != nil {
		status := 500
		if strings.HasPrefix(err.Error(), "invalid mimetype") {
			status = 400
		}
		shared.LogError(err)
		shared.WriteError(w, status, err)
		return
	}
	w.WriteHeader(201)
	w.Write([]byte(filename))
}

func FetchBookFile(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("name")
	obj, err := service.FetchBook(filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	defer obj.Close()

	stat, err := obj.Stat()
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	http.ServeContent(w, r, filename, stat.LastModified, obj)
}

func FetchBookFileInfo(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("name")
	obj, err := service.FetchBookStat(filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}

	json, _ := json.Marshal(obj)
	w.Write(json)
}

func DeleteBookFile(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("name")
	obj, err := service.FetchBook(filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	defer obj.Close()

	err = service.DeleteBook(filename, r)
	if err != nil {
		shared.LogError(err)
		shared.WriteError(w, 500, err)
		return
	}
	w.WriteHeader(204)
}
