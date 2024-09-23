package config


import (
	"github.com/caarlos0/env/v6"
	"log"
    "github.com/joho/godotenv"
 
)
type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT"`
	DBHost string `env:"DB_HOST"`
	DBName string `env:"DB_NAME"`
	DBUser string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBSSLMode string `env:"DB_SSLMODE"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil{
		log.Fatalf("Unable to load the .env: %v", err)
	}
	config := &EnvConfig{}

	if err := env.Parse(config); err !=nil{
		log.Fatalf("Unable to load variables from the .env: %v", err)
	}
	return config
}