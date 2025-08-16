package catalog

import "time"

type Author struct {
	ID          int
	Name        string
	Description string
	PhotoURL    string
}

type Book struct {
	ID           string
	FilePath     string
	Title        string
	Description  string
	Format       string
	CategoryID   int
	LanguageCode string
	Published    time.Time
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagToBook struct {
	TagID  int    `json:"tag_id"`
	BookID string `json:"book_id"`
}

type AuthorToBook struct {
	AuthorID int    `json:"author_id"`
	BookID   string `json:"book_id"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
