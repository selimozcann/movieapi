package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT            = getEnv("PORT", "")
	DB_TYPE         = getEnv("DB_TYPE", "")
	MONGO_DB_NAME   = getEnv("MONGO_DB_NAME", "")
	MONGO_URL_START = getEnv("MONGO_URL_START", "")
	MONGO_DB_URL    = getEnv("MONGO_DB_URL", "")
	MONGO_USER_NAME = getEnv("MONGO_USER_NAME", "")
	MONGO_PASSWORD  = getEnv("MONGO_PASSWORD", "")
)

func getEnv(name string, fallback string) string {
	godotenv.Load()

	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	log.Println(fmt.Sprintf(`Environment variable not found :: %v`, name))
	return "OK"
}
