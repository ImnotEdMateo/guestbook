package config

import (
	"os"
	"strings"
)

type Config struct {
	Port           string
	AllowedOrigins []string
}

func Load() *Config {
	port := os.Getenv("GUESTBOOK_PORT")
	if port == "" {
		port = "3000"
	}

	origins := os.Getenv("GUESTBOOK_ALLOWED_ORIGINS")
	if origins == "" {
		origins = "http://localhost"
	}

	return &Config{
		Port:           port,
		AllowedOrigins: strings.Split(origins, ","),
	}
}
