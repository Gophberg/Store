package usecase

import "net/http"

type ads interface {
	CreateAd(w http.ResponseWriter, r *http.Request)
	GetAllAds(w http.ResponseWriter, r *http.Request)
	GetAd(w http.ResponseWriter, r *http.Request)
}

type UseCase struct {
	ads ads
}

func NewUseCase(ads ads) *UseCase {
	return &UseCase{
		ads: ads,
	}
}
