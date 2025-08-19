package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	data, err := os.ReadFile("./test.pdf")
	if err != nil {
		log.Fatalln(err)
	}

	r, err := http.Post("http://localhost:8080/catalog/v1/files", "application/pdf", bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	fmt.Println(string(body))
}
