package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"strconv"
	"time"
)

type Config struct {
	Dbhost       string        `env:"POSTGRES_HOST"`
	Dbname       string        `env:"POSTGRES_DB_NAME"`
	Dbusername   string        `env:"POSTGRES_USER"`
	Dbpassword   string        `env:"POSTGRES_PASSWORD"`
	Dockerdbport int           `env:"DOCKER_DB_PORT"`
	MaxPoolSize  int           `env:"MAX_POOL_SIZE"`
	ConnAttempts int           `env:"CONN_ATTEMPTS"`
	ConnTimeout  time.Duration `env:"CONN_TIMEOUT"`
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
	config.Dockerdbport, err = strconv.Atoi(read["DOCKER_DB_PORT"])
	if err != nil {
		return config, fmt.Errorf("error reading DOCKER_DB_PORT from .env file %w", err)
	}
	config.MaxPoolSize, err = strconv.Atoi(read["MAX_POOL_SIZE"])
	if err != nil {
		return config, fmt.Errorf("error reading MAX_POOL_SIZE from .env file %w", err)
	}
	config.ConnAttempts, err = strconv.Atoi(read["CONN_ATTEMPTS"])
	if err != nil {
		return config, fmt.Errorf("error reading CONN_ATTEMPTS from .env file %w", err)
	}
	connTimeout, err := strconv.Atoi(read["CONN_TIMEOUT"])
	if err != nil {
		return config, fmt.Errorf("error reading CONN_TIMEOUT from .env file %w", err)
	}
	config.ConnTimeout = time.Duration(connTimeout) * time.Second

	return config, nil
}
