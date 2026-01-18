package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string

	PBAdminEmail    string
	PBAdminPassword string
	PBDataDir       string

	StreamAPIKey    string
	StreamAPISecret string
	StreamAppID     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	pbDataDir, err := getEnvRequired("PB_DATA_DIR")
	if err != nil {
		return nil, err
	}

	pbAdminEmail, err := getEnvRequired("PB_ADMIN_EMAIL")
	if err != nil {
		return nil, err
	}

	pbAdminPassword, err := getEnvRequired("PB_ADMIN_PASSWORD")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:            getEnvDefault("PORT", "3000"),
		LogLevel:        getEnvDefault("LOG_LEVEL", "info"),
		PBAdminEmail:    pbAdminEmail,
		PBAdminPassword: pbAdminPassword,
		PBDataDir:       pbDataDir,
		StreamAPIKey:    getEnvDefault("STREAM_API_KEY", ""),
		StreamAPISecret: getEnvDefault("STREAM_API_SECRET", ""),
		StreamAppID:     getEnvDefault("STREAM_APP_ID", ""),
	}

	return cfg, nil
}
