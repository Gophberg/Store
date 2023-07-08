package usecase

import "net/http"

type ads interface {
	createAd(w http.ResponseWriter, r *http.Request)
	getAllAds(w http.ResponseWriter, r *http.Request)
	getAd(w http.ResponseWriter, r *http.Request)
}

type UseCase struct {
	ads ads
}

func NewUseCase(ads ads) *UseCase {
	return &UseCase{
		ads: ads,
	}
}
