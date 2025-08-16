package catalog

import "time"

type Author struct {
	ID          int
	Name        string
	Description string
	PhotoURL    string
}

type Book struct {
	Hash         string
	FilePath         string
	Title        string
	Description  string
	Format       string
	CategoryID   int
	LanguageCode string
	Published    time.Time
}

type Category struct {
	ID   int
	Name string
}

type Tag struct {
	ID   int
	Name string
}

type TagToBook struct {
	TagID    int
	BookHash string
}

type AuthorToBook struct {
	AuthorID int
	BookHash string
}

type Language struct {
	Code string
	Name string
}
