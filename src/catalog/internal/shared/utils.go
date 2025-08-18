package shared

import (
	"log"
	"reflect"
)

func LogError(err error, data ...any) {
	log.Println("\033[0;31m[ERROR]\033[0m", err, "| data:", reflect.TypeOf(err).String(), data)
}
