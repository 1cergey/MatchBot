package config

import (
	"os"
)

type Config struct {
	TelegramToken string
	WebHookURL    string
	Port          string
}

func New() *Config {
	return &Config{
		TelegramToken: getEnv("telegramToken", ""),
		WebHookURL:    getEnv("webHookURL", ""),
		Port:          getEnv("port", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return defaultVal

}
