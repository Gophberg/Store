package Store

import (
	"github.com/shopspring/decimal"
	"time"
)

type Config struct {
	Dbhost       string `env:"POSTGRES_HOST"`
	Dbname       string `env:"POSTGRES_DB_NAME"`
	Dbusername   string `env:"POSTGRES_USER"`
	Dbpassword   string `env:"POSTGRES_PASSWORD"`
	Dockerdbport string `env:"DOCKER_DB_PORT"`
}

type Ad struct {
	Photo        []string        `json:"photo"`
	Id           int64           `json:"id"`
	Title        string          `json:"title"`
	Content      string          `json:"content"`
	Price        decimal.Decimal `json:"price"`
	CreationDate time.Time       `json:"datecreated"`
}

type Result struct {
	Id     int64
	Status bool
	Reason string
}

type QueryCredentials struct {
	OrderBy   string `json:"orderby"`
	Direction string `json:"direction"`
	Limit     string `json:"limit"`
	Offset    int    `json:"offset"`
}
