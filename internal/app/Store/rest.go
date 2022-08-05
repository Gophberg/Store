package Store

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
)

func (a Ad) createAd(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested create ad: %s\n", r.URL.Path)

	a.decodeData(w, r)

	// Validations
	if len(a.Content) > 10 {
		log.Println("[REST] The content is to large")
	}

	// Create new record
	id, err := a.createRecord(a)
	if err != nil {
		log.Println("[REST]", err)
	}
	log.Println("[REST] New ad created with id:", id)

	// Sending response
	js, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (a Ad) getAd(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested ad: %s\n", r.URL.Path)

	a.decodeData(w, r)

	ad, err := a.readRecord(a)
	if err != nil {
		log.Println("[REST]", err)
	}

	js, err := json.Marshal(ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (a Ad) getAllAds(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested all ads: %s\n", r.URL.Path)

	a.decodeData(w, r)

	ads, err := a.readRecords()
	if err != nil {
		log.Println("[REST]", err)
	}

	js, err := json.Marshal(ads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (a *Ad) decodeData(w http.ResponseWriter, r *http.Request) {
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

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[REST] Data to decode: %v\n", &a)
}
