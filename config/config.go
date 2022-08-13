package config

import (
	"io"
	"os"

	"github.com/joho/godotenv"
)

var (
	Config = loadConfig()

	DefaultWriter io.Writer = os.Stdout
)

type config struct {
	Host       string
	Port       string
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

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	c.Host = os.Getenv("SERVER_HOST")
	c.Port = os.Getenv("PORT")
	c.DbEngine = os.Getenv("DB_ENGINE")
	c.DbHost = os.Getenv("DB_HOST")
	c.DbName = os.Getenv("DB_NAME")
	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASS")
	c.JwtSecret = os.Getenv("APP_JWT_SECRET")
	c.APIBaseURL = os.Getenv("API_BASE_URL")

	return c
}
