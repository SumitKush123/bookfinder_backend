package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port                 string `env:"PORT"`
	Dburl                string `env:"DB_URL"`
	DbHost               string `env:"DB_HOST"`
	DbPort               string `env:"DB_PORT"`
	DbName               string `env:"DB_NAME"`
	DbUser               string `env:"DB_USER"`
	DbPassword           string `env:"DB_PASSWORD"`
	JWTSecret            string `env:"JWT_SECRET"`
	PrimaryEmail         string `env:"PRIMARY_EMAIL"`
	PrimaryEmailPassword string `env:"PRIMARY_EMAIL_PASSWORD"`
}

func GetConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Error loading .env file:", err)
	}

	return Configuration{
		Port:                 getEnv("PORT", "8001"),
		DbHost:               getEnv("DB_HOST", "localhost"),
		DbPort:               getEnv("DB_PORT", "5432"),
		DbName:               getEnv("DB_NAME", "book_finder_db"),
		DbUser:               getEnv("DB_USER", "book_finder_user"),
		DbPassword:           getEnv("DB_PASSWORD", "book_finder_pass123"),
		Dburl:                getEnv("DB_URL", ""),
		JWTSecret:            getEnv("JWT_SECRET", "book_finder_jwt_secret_key123"),
		PrimaryEmail:         getEnv("PRIMARY_EMAIL", ""),
		PrimaryEmailPassword: getEnv("PRIMARY_EMAIL_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
