package config

import (
	"os"
)

type Config struct {
	TelegramToken string
}

func New() *Config {
	return &Config{
		TelegramToken: getEnv("telegramToken", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return defaultVal

}
