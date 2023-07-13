package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gophberg/Store/internal/app/Store"
	"github.com/Gophberg/Store/pkg/postgres"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
	"time"
)

type StoreRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *StoreRepo {
	return &StoreRepo{
		Postgres: pg,
	}
}

func (a *Store.Ad) createRecord(c Store.Ad) (int64, error) {
	log.Printf("[DB] Reseived createRecord Credentials: %v", c)

	conn, err := connDB()
	if err != nil {
		return 0, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Println("[DB defer conn]", err)
		}
	}(conn, context.Background())

	c.CreationDate = time.Now()

	err = conn.QueryRow(context.Background(),
		`INSERT INTO store (title, content, photo, price, createdate) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		c.Title,
		c.Content,
		c.Photo,
		c.Price,
		c.CreationDate,
	).Scan(&a.Id)
	return a.Id, err
}

func (a Store.Ad) readRecord(c Store.Ad) (Store.Ad, error) {
	log.Printf("[DB] Requested ad with id: <%v>\n", c.Id)

	conn, err := connDB()
	if err != nil {
		return a, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Println("[DB defer conn]", err)
		}
	}(conn, context.Background())

	query := fmt.Sprintf(`SELECT title, content, photo, price, createdate FROM store WHERE id = %d;`, c.Id)
	if err := conn.QueryRow(context.Background(), query).Scan(
		&a.Title,
		&a.Content,
		&a.Photo,
		&a.Price,
		&a.CreationDate,
	); err != nil {
		if err == sql.ErrNoRows {
			return a, err
		}
	}
	return a, nil
}

func (a Store.Ad) readRecords(qc Store.QueryCredentials) ([]Store.Ad, error) {
	log.Printf("[DB] Requested all ads with credentials: <%v>\n", qc)
	conn, err := connDB()
	if err != nil {
		return nil, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Println("[DB defer conn]", err)
		}
	}(conn, context.Background())

	var ads []Store.Ad
	var f string

	if qc.Required {
		fields := strings.Fields(qc.Fields)
		f = strings.Join(fields, ",")
	} else {
		f = fmt.Sprint("*")
	}

	log.Println("[DB qc]", qc)
	log.Println("[DB f]", f)

	query := `SELECT $1 FROM "store" ORDER BY $2, $3 LIMIT $4 OFFSET $5`
	log.Println("[DB query]", query)
	rows, err := conn.Query(context.Background(), query, f, qc.OrderBy, qc.Direction, qc.Limit, qc.Offset)
	if err != nil {
		log.Println("[DB query]", err)
		return nil, err
	}
	for rows.Next() {
		var i Store.Ad
		if err := rows.Scan(
			&i.Id,
			&i.Title,
			&i.Content,
			&i.Photo,
			&i.Price,
			&i.CreationDate,
		); err != nil {
			log.Println("[DB loop]", err)
			return nil, err
		}
		ads = append(ads, i)
	}
	return ads, nil
}
