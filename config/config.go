// config/config.go
package config

import (
    "os"
)

type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
}

func GetConfig() *Config {
    return &Config{
        Port:        getEnv("PORT", "3000"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        JWTSecret:   getEnv("JWT_SECRET", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
