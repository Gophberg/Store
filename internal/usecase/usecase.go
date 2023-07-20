package usecase

import (
	"context"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
)

type AdsUseCase struct {
	repo Repository
}

func NewUseCase(r Repository) *AdsUseCase {
	return &AdsUseCase{
		repo: r,
	}
}

func (u *AdsUseCase) CreateAd(ctx context.Context, a entity.Ad) (uint64, error) {
	id, err := u.repo.CreateRecord(ctx, a)
	if err != nil {
		return 0, fmt.Errorf("AdsUseCase - CreateAd - u.repo.CreateRecord: %w", err)
	}

	return id, nil
}

func (u *AdsUseCase) GetAllAds(ctx context.Context, qc entity.QueryCredentials) ([]entity.Ad, error) {
	ad, err := u.repo.ReadRecords(ctx, qc)
	if err != nil {
		return []entity.Ad{}, fmt.Errorf("AdsUseCase - GetAd - u.repo.GetAd: %w", err)
	}

	return ad, nil
}

func (u *AdsUseCase) GetAd(ctx context.Context, a entity.Ad) (entity.Ad, error) {
	ad, err := u.repo.ReadRecord(ctx, a.Id)
	if err != nil {
		return entity.Ad{}, fmt.Errorf("AdsUseCase - GetAd - u.repo.GetAd: %w", err)
	}

	return ad, nil
}
