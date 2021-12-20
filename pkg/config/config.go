package config

import (
	"os"
	"strconv"
)

const (
	DefaultPort = 5432
)

type Config struct {
	DbConfig *DbConfig
}

type DbConfig struct {
	Url  string
	Host string
	Port int
	User string
	Pass string
	Name string
}

func getDbConfig() *DbConfig {
	return &DbConfig{
		Url:  getEnv("DB_URL", ""),
		Host: getEnv("DB_HOST", ""),
		Port: getIntEnv("DB_PORT", DefaultPort),
		User: getEnv("DB_USER", ""),
		Pass: getEnv("DB_PASS", ""),
		Name: getEnv("DB_NAME", ""),
	}
}

func New() *Config {
	return &Config{
		DbConfig: getDbConfig(),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		val, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return defaultVal
		}

		return int(val)
	}

	return defaultVal
}
