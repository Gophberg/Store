package usecase

import (
	"context"
	"github.com/Gophberg/Store/internal/entity"
)

type Repository interface {
	CreateRecord(context.Context, entity.Ad) (uint64, error)
	ReadRecords(context.Context, entity.QueryCredentials) ([]entity.Ad, error)
	ReadRecord(context.Context, uint64) (entity.Ad, error)
}
