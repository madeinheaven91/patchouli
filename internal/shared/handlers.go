// Package shared
//
// For shared stuff
package shared

import "net/http"

// HealthHandler ...
//
// I dont think i need it right now, but why not
type HealthHandler struct{}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UP\n"))
}
