package config

import "os"

type envConfig struct {
	ZincUser string
	ZincPass string
	ZincURL  string
}

var EnvVars envConfig

func loadEnvVars() envConfig {
	return envConfig{
		ZincUser: os.Getenv("ZINC_FIRST_ADMIN_USER"),
		ZincPass: os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"),
		ZincURL:  os.Getenv("ZINC_SERVER_URL"),
	}
}

func init() {
	EnvVars = loadEnvVars()
}
