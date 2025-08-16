package shared

import (
	"log"
	"net/http"
	"time"
)

func LoggingMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		end := time.Now()
		log.Printf("%s %s %s | sent %d | handled in %d ms",
			r.RemoteAddr,
			r.Method,
			r.URL.String(),
			r.ContentLength,
			end.Sub(start).Milliseconds())
	}
}

// TODO: integrate with auth service
func AuthMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requesting auth for", r.URL.String())
		next(w, r)
	}
}
