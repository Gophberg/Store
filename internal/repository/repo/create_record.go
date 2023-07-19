package repo

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

func (s *StoreRepo) CreateRecord(ctx context.Context, a entity.Ad) (uint64, error) {

	//a.CreationDate = time.Now()

	sql, args, err := s.Builder.
		Insert("store").
		Columns("title", "content", "photo", "price", "createdate").
		Values(a.Title, a.Content, a.Photo, a.Price, a.CreationDate).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("StoreRepo - CreateRecord - s.Builder: %w", err)
	}

	err = s.Pool.QueryRow(ctx, sql, args...).Scan(&a.Id)
	if err != nil {
		return 0, fmt.Errorf("StoreRepo - CreateRecord - s.Pool.Exec: %w", err)
	}

	return a.Id, nil
}
