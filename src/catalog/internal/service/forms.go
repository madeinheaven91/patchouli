package service

type BookPostForm struct {
	Title       string
	Author      string
	Description string
	Format      string
	CategoryID  int
	Document    string
}

type AuthorPostForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url,omitempty"`
}

type RequestPostForm struct {
	Filename     string   `json:"filename"`
	Title        string   `json:"title"`
	AuthorName   string   `json:"author_name"`
	Description  string   `json:"description"`
	Category     string   `json:"category"`
	LanguageCode string   `json:"language_code"`
	Tags         []string `json:"tags"`
}
