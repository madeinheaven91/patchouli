// Package models
package models

import "time"

type Author struct {
	ID          string
	Name        string
	Description string
	PhotoURL    string
}

type Book struct {
	ID           string
	Filename     string
	Title        string
	AuthorID     string
	Description  string
	Format       string
	Category     string
	LanguageCode string
	Published    time.Time
}

type Request struct {
	ID           string
	Filename     string
	Title        string
	AuthorName   string
	Description  string
	Category     string
	LanguageCode string
	Added        time.Time
}

type Tag struct {
	Name string
}

func (t Tag) String() string {
	return t.Name
}

type TagToBook struct {
	TagName string `json:"tag_name"`
	BookID  string `json:"book_id"`
}

type TagToRequest struct {
	TagName   string `json:"tag_name"`
	RequestID string `json:"request_id"`
}

type AuthorToBook struct {
	AuthorID int    `json:"author_id"`
	BookID   string `json:"book_id"`
}
