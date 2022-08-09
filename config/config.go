package config

import (
	"io"
	"os"
)

var (
	Config = loadConfig()

	DefaultWriter io.Writer = os.Stdout
)

type config struct {
	Host       string
	DbEngine   string
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
	APIBaseURL string
}

func loadConfig() *config {
	c := new(config)

	c.Host = os.Getenv("SERVER_HOST")
	c.DbEngine = os.Getenv("DB_ENGINE")
	c.DbHost = os.Getenv("DB_HOST")
	c.DbName = os.Getenv("DB_NAME")
	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASS")
	c.JwtSecret = os.Getenv("APP_JWT_SECRET")
	c.APIBaseURL = os.Getenv("API_BASE_URL")

	return c
}
