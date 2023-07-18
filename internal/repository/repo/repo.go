package repo

import (
	"github.com/Gophberg/Store/pkg/postgres"
)

type StoreRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *StoreRepo {
	return &StoreRepo{
		Postgres: pg,
	}
}
