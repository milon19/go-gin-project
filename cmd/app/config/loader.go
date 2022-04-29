package config

import (
	"os"
)

var (
	GinPort      = ""
	DatabaseName = ""
	MongoSSL     = ""
	MongoUrl     = ""
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
	DatabaseName = os.Getenv("DATABASE_NAME")
	MongoSSL = GetEnvDefault("MONGO_SSL", "false")
	MongoUrl = os.Getenv("MONGO_URL")
}
