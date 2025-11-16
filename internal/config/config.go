package config

import (
	"os"
	"time"
)

type Config struct {
	ServerAddr   string
	DBUrl        string
	AutoLockTime time.Duration
	JwtSecret    string
}

func Load() *Config {
	return &Config{
		ServerAddr:   getEnv("SERVER_ADDR", ":8080"),
		DBUrl:        getEnv("DATABASE_URL", "./data/"),
		JwtSecret:    getEnv("JWT_SECRET", "default-secret"),
		AutoLockTime: getDurationEnv("AUTO_LOCK_TIME", 2*time.Hour),
	}
}
func getDurationEnv(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
