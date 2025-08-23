package shared

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func LogError(err error, data ...any) {
	log.Println("\033[0;31m[ERROR]\033[0m", err, "| data:", reflect.TypeOf(err).String(), data)
}

func WriteError(w http.ResponseWriter, statusCode int, message any) {
	var msg string
	switch v := message.(type) {
	case string:
		msg = v
	case error:
		msg = v.Error()
	case fmt.Stringer:
		msg = v.String()
	default:
		msg = fmt.Sprintf("%v", v)
	}
	response := fmt.Sprintf(`{"status_code":"%d","message":"%v"}`, statusCode, strings.ReplaceAll(msg, "\"", "'"))
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

func ToFilename(str string) string {
	r := strings.NewReplacer(
		" ", "_",
		".", "",
		"/", "",
		`\`, "",
		"<", "",
		">", "",
		":", "",
		`"`, "",
		"|", "",
		"?", "",
		"*", "",
	)
	return strings.ToLower(r.Replace(str))
}
