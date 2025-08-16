// Package shared
//
// Shared stuff
package shared

import "os"

var Config Cfg

type Cfg struct {
	DBUser string
	DBPass string
	DBPort string
}

func InitFromEnv() {
	Config.DBUser = os.Getenv("DB_USER")
	Config.DBPass = os.Getenv("DB_PASS")
	Config.DBPort = os.Getenv("DB_PORT")
}
