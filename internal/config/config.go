package config

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	ServerPort     string
	MetricsPort    string
	LogLevel       slog.Level
	AgifyURL       string
	GenderizeURL   string
	NationalizeURL string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "effective_mobile"),
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		MetricsPort:    getEnv("METRICS_PORT", "9090"),
		AgifyURL:       getEnv("AGIFY_URL", "https://api.agify.io"),
		GenderizeURL:   getEnv("GENDERIZE_URL", "https://api.genderize.io"),
		NationalizeURL: getEnv("NATIONALIZE_URL", "https://api.nationalize.io"),
	}

	logLevel := getEnv("LOG_LEVEL", "info")

	switch logLevel {
	case "debug":
		cfg.LogLevel = slog.LevelDebug
	case "warn":
		cfg.LogLevel = slog.LevelWarn
	case "error":
		cfg.LogLevel = slog.LevelError
	default:
		cfg.LogLevel = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(cfg.LogLevel)
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
