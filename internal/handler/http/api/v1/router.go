package v1

import (
	"github.com/gorilla/mux"
)

func (h *Handler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/ads", h.CreateAd).Methods("POST")
	r.HandleFunc("/ads", h.GetAllAds).Methods("GET")
	r.HandleFunc("/ads/{id}", h.GetAd).Methods("GET")
}
