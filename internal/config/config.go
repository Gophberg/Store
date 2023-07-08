package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

type Config struct {
	Dbhost       string `env:"POSTGRES_HOST"`
	Dbname       string `env:"POSTGRES_DB_NAME"`
	Dbusername   string `env:"POSTGRES_USER"`
	Dbpassword   string `env:"POSTGRES_PASSWORD"`
	Dockerdbport string `env:"DOCKER_DB_PORT"`
}

func NewConfig() (Config, error) {

	var config Config

	err := godotenv.Load()
	if err != nil {
		return config, fmt.Errorf("error loading .env file %w", err)
	}

	read, err := godotenv.Read()
	if err != nil {
		return config, fmt.Errorf("error reading .env file %w", err)
	}

	config.Dbhost = read["POSTGRES_HOST"]
	config.Dbname = read["POSTGRES_DB_NAME"]
	config.Dbusername = read["POSTGRES_USER"]
	config.Dbpassword = read["POSTGRES_PASSWORD"]
	config.Dockerdbport = read["DOCKER_DB_PORT"]

	return config, nil
}
