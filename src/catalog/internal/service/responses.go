package service

import (
	"time"

	"catalog/internal/models"
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

func RequestResponseFromModel[T string | models.Tag](model models.Request, tags []T) RequestResponse {
	return RequestResponse{
		ID:           model.ID,
		Filename:     model.Filename,
		Title:        model.Title,
		AuthorName:   model.AuthorName,
		Description:  model.Description,
		Category:     model.Category,
		LanguageCode: model.LanguageCode,
		Added:        model.Added,
		Tags:         mapTags(tags),
	}
}

func mapTags[T string | models.Tag](items []T) []string {
	// FIXME: looks a bit ugly
	res := make([]string, len(items))
	for i := range items {
		switch t := any(items[i]).(type) {
		case string:
			res[i] = t
		case models.Tag:
			res[i] = t.Name
		}
	}
	return res
}

type BookResponse struct {
	ID           string
	Filename     string
	Title        string
	AuthorName   string
	Description  string
	Category     string
	LanguageCode string
	Published    time.Time
	Tags         []string
}

func BookResponseFromModel[T string | models.Tag](model models.Book, tags []T) BookResponse {
	return BookResponse{
		ID:           model.ID,
		Filename:     model.Filename,
		Title:        model.Title,
		AuthorName:   model.AuthorName,
		Description:  model.Description,
		Category:     model.Category,
		LanguageCode: model.LanguageCode,
		Published:    model.Published,
		Tags:         mapTags(tags),
	}
}
