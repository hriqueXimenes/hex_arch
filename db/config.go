package db

import (
	"os"
)

// Config contains the configurations for database.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// InitConfig load the configuration for database.
func InitConfig() *Config {
	config := &Config{
		Host:     getEnv("database_host", "localhost"),
		Port:     getEnv("database_port", "5431"),
		User:     getEnv("database_user", "postgres"),
		Password: getEnv("database_password", ""),
		DbName:   getEnv("database_dbname", "postgres"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
