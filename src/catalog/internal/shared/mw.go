package shared

import (
	"log"
	"net/http"
)

// TODO: integrate with auth service
func AuthMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requesting auth for", r.URL.String())
		next(w, r)
	}
}
