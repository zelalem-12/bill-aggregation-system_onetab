package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func (config *Config) Validate() error {
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return err
	}
	return nil
}

type Config struct {
	SERVER_HOST  string `validate:"required"`
	SERVER_PORT  int    `validate:"required"`
	FRONTEND_URL string `validate:"required"`

	ACCESS_TOKEN_KEY  string `validate:"required"`
	REFRESH_TOKEN_KEY string `validate:"required"`

	ACCESS_TOKEN_EXPIRY  string `validate:"required"`
	REFRESH_TOKEN_EXPIRY string `validate:"required"`

	SMTP_HOST       string `validate:"required"`
	SMTP_PORT       int    `validate:"required"`
	SENDER_EMAIL    string `validate:"required"`
	SENDER_PASSWORD string `validate:"required"`

	POSTGRES_HOST     string `validate:"required"`
	POSTGRES_PORT     int    `validate:"required"`
	POSTGRES_DATABASE string `validate:"required"`
	POSTGRES_USER     string `validate:"required"`
	POSTGRES_PASSWORD string `validate:"required"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	serverPort, err := strconv.Atoi(strings.TrimSpace(os.Getenv("SERVER_PORT")))
	if err != nil {
		return nil, err
	}

	postGresPort, err := strconv.Atoi(strings.TrimSpace(os.Getenv("POSTGRES_PORT")))
	if err != nil {
		return nil, err
	}
	smtpPort, err := strconv.Atoi(strings.TrimSpace(os.Getenv("SMTP_PORT")))
	if err != nil {
		return nil, err
	}

	config := Config{
		SERVER_HOST:  strings.TrimSpace(os.Getenv("SERVER_HOST")),
		SERVER_PORT:  serverPort,
		FRONTEND_URL: strings.TrimSpace(os.Getenv("FRONTEND_URL")),

		ACCESS_TOKEN_KEY:  strings.TrimSpace(os.Getenv("ACCESS_TOKEN_KEY")),
		REFRESH_TOKEN_KEY: strings.TrimSpace(os.Getenv("REFRESH_TOKEN_KEY")),

		ACCESS_TOKEN_EXPIRY:  strings.TrimSpace(os.Getenv("ACCESS_TOKEN_EXPIRY")),
		REFRESH_TOKEN_EXPIRY: strings.TrimSpace(os.Getenv("REFRESH_TOKEN_EXPIRY")),

		SMTP_HOST:       strings.TrimSpace(os.Getenv("SMTP_HOST")),
		SMTP_PORT:       smtpPort,
		SENDER_EMAIL:    strings.TrimSpace(os.Getenv("SENDER_EMAIL")),
		SENDER_PASSWORD: strings.TrimSpace(os.Getenv("SENDER_PASSWORD")),

		POSTGRES_HOST:     strings.TrimSpace(os.Getenv("POSTGRES_HOST")),
		POSTGRES_PORT:     postGresPort,
		POSTGRES_DATABASE: strings.TrimSpace(os.Getenv("POSTGRES_DATABASE")),
		POSTGRES_USER:     strings.TrimSpace(os.Getenv("POSTGRES_USER")),
		POSTGRES_PASSWORD: strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD")),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil

}
