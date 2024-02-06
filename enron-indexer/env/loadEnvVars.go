package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	ZincUser string
	ZincPass string
	ZincHost string
	ZincPort string
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
		ZincHost: os.Getenv("ZINC_SERVER_ADDRESS"),
		ZincPort: os.Getenv("ZINC_SERVER_PORT"),
	}
}

func init() {
	EnvVars = loadEnvVars()
}
