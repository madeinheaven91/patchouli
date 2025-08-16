package shared

import "log"

func LogError(err error, data ...any) {
	log.Println("\033[0;31m[ERROR]\033[0m", err)
	log.Println("Error data:", data)
}
