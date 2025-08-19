package handlers

import (
	"encoding/json"
	"net/http"

	"catalog/internal/service"
	"catalog/internal/shared"
)

// NOTE: im not sure whether multipart upload is required, maybe ill implement it later
// for now i will use a simple solution. Maybe its bad, idk

// UploadBookFile handles a file upload in single http request.
func UploadBookFile(w http.ResponseWriter, r *http.Request) {
	filename, err := service.UploadBook(r)
	if err != nil {
		// FIXME: it can be not only 500, depends on UploadBook output
		shared.LogError(err)
		w.WriteHeader(500)
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
		w.WriteHeader(500)
		return
	}
	defer obj.Close()

	stat, err := obj.Stat()
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}

	http.ServeContent(w, r, filename, stat.LastModified, obj)
}

func FetchBookFileInfo(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("name")
	obj, err := service.FetchBookStat(filename, r)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
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
		w.WriteHeader(500)
		return
	}
	defer obj.Close()

	err = service.DeleteBook(filename, r)
	if err != nil {
		shared.LogError(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
