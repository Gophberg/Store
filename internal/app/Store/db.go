package Store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"time"
)

func connDB() (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.Dbusername, config.Dbpassword, config.Dbhost, config.Dockerdbport, config.Dbname)
	log.Println("[DB] uri", url)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	conn.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &shopspring.Numeric{},
		Name:  "numeric",
		OID:   pgtype.NumericOID,
	})

	return conn, err
}

func (a *Ad) createRecord(c Ad) (int64, error) {
	conn, err := connDB()
	if err != nil {
		return 0, err
	}

	log.Printf("[DB] Reseived createRecord Credentials: %v", c)

	c.CreationDate = time.Now()
	//c.CreationDate = time.Now().Format(time.RFC3339)

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

func (a Ad) readRecord(c Ad) (Ad, error) {
	conn, err := connDB()
	if err != nil {
		return a, err
	}

	log.Printf("[DB] Query request <%v> ad\n", c.Id)

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
	//log.Printf("[DB] Read record <%d>\n", a.Id)
	return a, nil
}

func (a Ad) readRecords(qc QueryCredentials) ([]Ad, error) {
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

	var ads []Ad

	//SQL := ` SELECT "id","price" FROM "store" ORDER BY "id" LIMIT $2 OFFSET $1`

	query := `SELECT * FROM "store" ORDER BY $1, $2 LIMIT $3 OFFSET $4`
	//query := fmt.Sprintf(`SELECT * FROM store LIMIT 5 OFFSET 5;`)
	rows, err := conn.Query(context.Background(), query, qc.By, qc.Order, qc.Limit, qc.Offset)
	if err != nil {
		log.Println("[DB query]", err)
		return nil, err
	}
	for rows.Next() {
		var i Ad
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
