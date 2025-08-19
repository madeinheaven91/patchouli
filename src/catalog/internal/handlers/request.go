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
