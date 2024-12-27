package config

import (
	"log"
    "os"

    "github.com/joho/godotenv"
)

// Init loads environment variables
func Init() {
	os.Setenv("TZ", "UTC")

	mode := os.Getenv("GIN_MODE")

	envFile := ".env.development"
	if mode == "release" {
		envFile = ".env.production"
	}

	err := godotenv.Load(envFile)
	if err != nil {
	  log.Println("Error loading " + envFile + " file")
	}
}
