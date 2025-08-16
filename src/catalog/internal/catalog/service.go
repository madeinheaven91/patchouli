package catalog

import (
	"encoding/base64"
	"os"
)

type BookPostForm struct {
	Title       string
	Author      string
	Description string
	Format      string
	CategoryID  int
	Document    string
}

func FetchBookDocument(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func DecodeDocument(document string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(document)
}

func EncodeDocument(document []byte) string {
	return base64.StdEncoding.EncodeToString(document)
}

type AuthorPostForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url,omitempty"`
}
