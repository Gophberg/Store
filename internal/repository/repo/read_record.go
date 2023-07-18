package repo

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

func (s *StoreRepo) readRecord(ctx context.Context, c entity.Ad) (entity.Ad, error) {

	sql, _, err := s.Builder.
		Select("id", "title", "content", "photo", "price", "createdate").
		From("store").
		Where("id = ?", c.Id).
		ToSql()
	if err != nil {
		return entity.Ad{}, fmt.Errorf("StoreRepo - readRecord - s.Builder: %w", err)
	}

	if err = s.Pool.QueryRow(ctx, sql).Scan(
		&c.Id,
		&c.Title,
		&c.Content,
		&c.Photo,
		&c.Price,
		&c.CreationDate,
	); err != nil {
		return entity.Ad{}, fmt.Errorf("StoreRepo - readRecord - s.Pool.QueryRow: %w", err)
	}

	return c, nil
}
