package service

import (
	"time"

	"github.com/madeinheaven91/patchouli/internal/models"
)

type RequestResponse struct {
	ID           string
	Filename     string
	Title        string
	AuthorName   string
	Description  string
	Category     string
	LanguageCode string
	Added        time.Time
	Tags         []string
}

func RequestResponseFromModel(model models.Request, tags []string) RequestResponse {
	return RequestResponse{
		ID:           model.ID,
		Filename:     model.Filename,
		Title:        model.Title,
		AuthorName:   model.AuthorName,
		Description:  model.Description,
		Category:     model.Category,
		LanguageCode: model.LanguageCode,
		Added:        model.Added,
		Tags:         tags,
	}
}
