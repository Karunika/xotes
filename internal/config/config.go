package config

import (
	"os"

	godotenv "github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	host     string
	port     string
	name     string
	user     string
	password string
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
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
	}

	return config
}
