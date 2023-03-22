package config

import (
	"os"

	godotenv "github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func init() {
	err := godotenv.Load()

	if err != nil {
		return
	}
}

func GetServerConfig() ServerConfig {
	config := ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	return config
}

func GetDatabaseConfig() DatabaseConfig {
	config := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return config
}
