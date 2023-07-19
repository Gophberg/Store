package repo

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

func (s *StoreRepo) ReadRecords(ctx context.Context, qc entity.QueryCredentials) ([]entity.Ad, error) {

	sql, _, err := s.Builder.
		Select("id", "title", "content", "photo", "price", "createdate").
		From("store").OrderByClause(qc.OrderBy, qc.Direction).Limit(qc.Limit).Offset(qc.Offset).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("StoreRepo - ReadRecords - s.Builder: %w", err)
	}

	rows, err := s.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("StoreRepo - ReadRecords - s.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Ad, 0)

	for rows.Next() {
		e := entity.Ad{}

		if err := rows.Scan(
			&e.Id,
			&e.Title,
			&e.Content,
			&e.Photo,
			&e.Price,
			&e.CreationDate,
		); err != nil {
			return nil, fmt.Errorf("StoreRepo - ReadRecords - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil

	//var ads []entity.Ad
	//var f string
	//
	//if qc.Required {
	//	fields := strings.Fields(qc.Fields)
	//	f = strings.Join(fields, ",")
	//} else {
	//	f = fmt.Sprint("*")
	//}
	//
	//log.Println("[DB qc]", qc)
	//log.Println("[DB f]", f)
	//
	//query := `SELECT $1 FROM "store" ORDER BY $2, $3 LIMIT $4 OFFSET $5`
	//log.Println("[DB query]", query)
	//rows, err := conn.Query(context.Background(), query, f, qc.OrderBy, qc.Direction, qc.Limit, qc.Offset)
	//if err != nil {
	//	log.Println("[DB query]", err)
	//	return nil, err
	//}
	//for rows.Next() {
	//	var i entity.Ad
	//	if err := rows.Scan(
	//		&i.Id,
	//		&i.Title,
	//		&i.Content,
	//		&i.Photo,
	//		&i.Price,
	//		&i.CreationDate,
	//	); err != nil {
	//		log.Println("[DB loop]", err)
	//		return nil, err
	//	}
	//	ads = append(ads, i)
	//}
	//return ads, nil
}
