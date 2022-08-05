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

func (a Ad) readRecord(c Ad) (Ad, error) {
	db, err := ConnDB()
	if err != nil {
		return a, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	log.Printf("[DB] Query request <%v> ad\n", c.Id)

	query := fmt.Sprintf(`SELECT title, photo, price FROM store WHERE id = %d;`, c.Id)
	if err := db.QueryRow(query).Scan(
		&a.Title,
		&a.Photo,
		&a.Price,
	); err != nil {
		if err == sql.ErrNoRows {
			return a, err
		}
	}
	//log.Printf("[DB] Read record <%d>\n", a.Id)
	return a, nil
}

func (a Ad) readRecords() ([]Ad, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var ads []Ad

	query := fmt.Sprintf(`SELECT * FROM store;`)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i Ad
		if err := rows.Scan(
			&i.Id,
			&i.Title,
			&i.Photo,
			&i.Price,
		); err != nil {
			return nil, err
		}
		ads = append(ads, i)
	}
	return ads, nil
}
