package v1

import (
	"context"
	"github.com/Gophberg/Store/internal/entity"
)

type UseCase interface {
	CreateAd(ctx context.Context, a entity.Ad) (uint64, error)
	GetAllAds(ctx context.Context, qc entity.QueryCredentials) (interface{}, error)
	GetAd(ctx context.Context, a entity.Ad) (interface{}, error)
}

type Handler struct {
	uc UseCase
	//logger log.Logger
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}
