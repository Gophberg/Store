package repo

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

func (s *StoreRepo) ReadRecord(ctx context.Context, id uint64) (entity.Ad, error) {

	var ad entity.Ad

	sql, _, err := s.Builder.
		Select("id", "title", "content", "photo", "price", "createdate").
		From("store").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return ad, fmt.Errorf("StoreRepo - ReadRecord - s.Builder: %w", err)
	}

	if err = s.Pool.QueryRow(ctx, sql).Scan(
		&ad.Id,
		&ad.Title,
		&ad.Content,
		&ad.Photo,
		&ad.Price,
		&ad.CreationDate,
	); err != nil {
		return ad, fmt.Errorf("StoreRepo - ReadRecord - s.Pool.QueryRow: %w", err)
	}

	return ad, nil
}
