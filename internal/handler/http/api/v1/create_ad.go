package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Gophberg/Store/internal/entity"
	"log"
	"mime"
	"net/http"
)

func (h *Handler) CreateAd(w http.ResponseWriter, r *http.Request) {
	//log.Printf("[REST] Requested create ad: %s\n", r.URL.Path)

	var result entity.Result
	var a entity.Ad

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

	// Validations
	switch {
	case len(a.Title) > 200:
		log.Println("[REST] The title is to large")
		result.Reason = "The title is to large"
		result.Status = false
	case len(a.Content) > 1000:
		log.Println("[REST] The content is to large")
		result.Reason = "The content is to large"
		result.Status = false
	case len(a.Photo) > 3:
		log.Println("[REST] Limit of 3 photos exceeded")
		result.Reason = "Limit of 3 photos exceeded"
		result.Status = false
	default:
		// Create new record
		id, err := h.uc.CreateAd(r.Context(), a)
		if err != nil {
			log.Println("[REST]", fmt.Errorf("router - createRecord: %w", err))
			//return fmt.Errorf("router - createRecord: %w", err)
		} else {
			result.Id = id
			result.Reason = "Add created"
			result.Status = true
			log.Println("[REST] New ad created with id:", id)
		}
	}

	// Sending response
	js, err := json.Marshal(result)
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
