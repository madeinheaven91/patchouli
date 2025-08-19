// Package service
package service

import (
	"encoding/base64"
	"os"
)

func FetchBookDocument(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func DecodeDocument(document string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(document)
}

func EncodeDocument(document []byte) string {
	return base64.StdEncoding.EncodeToString(document)
}
