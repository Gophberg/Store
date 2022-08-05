package Store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnDB() (*sql.DB, error) {
	url := fmt.Sprintf("host=%v user=%v password='%v' dbname=%v sslmode=disable", config.Dbhost, config.Dbusername, config.Dbpassword, config.Dbname)
	log.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (a *Ad) createRecord(c Ad) (int64, error) {
	db, err := ConnDB()
	if err != nil {
		return 0, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	log.Printf("[DB] Reseived createRecord Credentials: %v", c)

	err = db.QueryRow(
		`INSERT INTO store (title, photo, price) 
		VALUES ($1, $2, $3) RETURNING id`,
		c.Title,
		c.Photo,
		c.Price,
	).Scan(&a.Id)
	return a.Id, err
}
