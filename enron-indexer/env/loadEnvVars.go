package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	ZincUser string
	ZincPass string
	ZincURL  string
}

var EnvVars envConfig

func loadEnvVars() envConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return envConfig{
		ZincUser: os.Getenv("ZINC_FIRST_ADMIN_USER"),
		ZincPass: os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"),
		ZincURL:  os.Getenv("ZINC_SERVER_URL"),
	}
}

func init() {
	EnvVars = loadEnvVars()
}
