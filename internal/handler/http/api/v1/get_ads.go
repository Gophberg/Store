package v1

import (
	"encoding/json"
	"github.com/Gophberg/Store/internal/entity"
	"mime"
	"net/http"
)

func (h *Handler) GetAllAds(w http.ResponseWriter, r *http.Request) {
	//log.Printf("[REST] Requested all ads: %s\n", r.URL.Path)

	var qc entity.QueryCredentials
	ctx := r.Context()

	contentType := r.Header.Get("Content-Type")

	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	// Decoding request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&qc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Get from repository
	ads, err := h.uc.GetAllAds(ctx, qc)

	// Sending response
	js, err := json.Marshal(ads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//write, err := w.Write(js)
	_, err = w.Write(js)
	if err != nil {
		return
	}
	//log.Printf("[REST] %v bytes written to ResponseWriter", write)
}
