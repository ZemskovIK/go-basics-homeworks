package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Не удалость найти env файл")
	}
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}, nil
}
