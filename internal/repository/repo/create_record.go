package repo

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

func (s *StoreRepo) createRecord(ctx context.Context, a entity.Ad) error {

	//a.CreationDate = time.Now()

	sql, args, err := s.Builder.
		Insert("store").
		Columns("title", "content", "photo", "price", "createdate").
		Values(a.Title, a.Content, a.Photo, a.Price, a.CreationDate).
		ToSql()
	if err != nil {
		return fmt.Errorf("StoreRepo - createRecord - s.Builder: %w", err)
	}

	_, err = s.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("StoreRepo - createRecord - s.Pool.Exec: %w", err)
	}

	return nil
}
