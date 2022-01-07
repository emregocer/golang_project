// package config provides the environment loading and management functions and structs.
package config

import (
	"os"
)

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Port     string
	Host     string
}

type Config struct {
	Database DatabaseConfig
	JWTKey   string
}

// New returns a new Config struct
func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Username: getEnv("PG_USER", ""),
			Password: getEnv("PG_PASS", ""),
			Host:     getEnv("PG_HOST", "localhost"),
			Name:     getEnv("DATABASE_NAME", ""),
			Port:     getEnv("DB_PORT", "5432"),
		},
		JWTKey: getEnv("JWT_KEY", ""),
	}
}

// getEnv returns the environment values or defaults for keys from the loaded environment.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
