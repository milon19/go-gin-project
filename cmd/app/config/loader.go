package config

import (
	"os"
)

var (
	GinPort = ""
)

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func SetEnvironment() {
	//Token = os.Getenv("TOKEN")
	GinPort = GetEnvDefault("GIN_PORT", "8080")
}
