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
	_ = godotenv.Load() // no falla si no existe

	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		HTTPPort:  getEnv("HTTP_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "FE2A2B346D1593EFAEF74A9AE5735"),
		DBURL:     getEnv("DATABASE_URL", "postgres://go:1234@localhost:5432/go_erp?sslmode=disable"),
	}

	if cfg.JWTSecret == "FE2A2B346D1593EFAEF74A9AE5735" {
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
