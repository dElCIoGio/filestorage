package config

import (
	"os"
)

type ENV string

const (
	DEV  ENV = "dev"
	PROD     = "prod"
	TEST     = "test"
)

type Config struct {
	Env  ENV
	Port string
}

func GetConfig() Config {
	return Config{
		Env:  ENV(get("APP_ENV", string(DEV))),
		Port: get("APP_PORT", "8080"),
	}
}

func get(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
