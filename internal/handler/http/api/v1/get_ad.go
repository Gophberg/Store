package v1

import (
	"encoding/json"
	"github.com/Gophberg/Store/internal/entity"
	"mime"
	"net/http"
)

func (h *Handler) GetAd(w http.ResponseWriter, r *http.Request) {
	//log.Printf("[REST] Requested ad: %s\n", r.URL.Path)

	var a entity.Ad
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
	if err := dec.Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Get from repository
	ad, err := h.uc.GetAd(ctx, a)

	// Sending response
	js, err := json.Marshal(ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	//write, err := w.Write(js)
	_, err = w.Write(js)
	if err != nil {
		return
	}
	//log.Printf("[REST] %v bytes written to ResponseWriter", write)
}
