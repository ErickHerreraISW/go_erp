package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	HTTPPort  string
	JWTSecret string
	DBURL     string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] .env file not found, relying on environment variables")
	}

	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		HTTPPort:  getEnv("HTTP_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "change-me"),
		DBURL:     getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable"),
	}

	if cfg.JWTSecret == "change-me" {
		log.Println("[WARN] JWT_SECRET usando valor por defecto. Cámbialo en producción.")
	}
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
