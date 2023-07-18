package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(config *Config) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  config.MaxPoolSize,
		connAttempts: config.ConnAttempts,
		connTimeout:  config.ConnTimeout,
	}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.Dbusername, config.Dbpassword, config.Dbhost, config.Dockerdbport, config.Dbname)
	//url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.Dbusername, config.Dbpassword, config.Dbhost, config.Dockerdbport, config.Dbname)

	dbpool, err := pgxpool.New(context.Background(), url)
	//dbpool, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w\n", err)
	}
	defer dbpool.Close()

	dbpool.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &shopspring.Numeric{},
		Name:  "numeric",
		OID:   pgtype.NumericOID,
	})

	return dbpool, nil
}
