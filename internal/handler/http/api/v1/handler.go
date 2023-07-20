package v1

import (
	"context"
	"github.com/Gophberg/Store/internal/entity"
)

type HandlerUseCase interface {
	CreateAd(ctx context.Context, a entity.Ad) (uint64, error)
	GetAllAds(ctx context.Context, qc entity.QueryCredentials) ([]entity.Ad, error)
	GetAd(ctx context.Context, a entity.Ad) (entity.Ad, error)
}

type Handler struct {
	uc HandlerUseCase
}

func NewHandler(uc HandlerUseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}
