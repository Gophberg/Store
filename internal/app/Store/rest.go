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

	var result Result

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
	default:
		// Create new record
		id, err := a.createRecord(a)
		result.Id = id
		result.Reason = "Add created"
		result.Status = true
		if err != nil {
			log.Println("[REST]", err)
		}
		log.Println("[REST] New ad created with id:", id)
	}

	// Sending response
	js, err := json.Marshal(result)
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

	var qc QueryCredentials

	qc.decodeData(w, r)

	log.Println("[REST qc]", qc)

	ads, err := a.readRecords(qc)
	if err != nil {
		log.Println("[REST from DB]", err)
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
	if err := dec.Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (qc *QueryCredentials) decodeData(w http.ResponseWriter, r *http.Request) {
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
	if err := dec.Decode(&qc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := dec.Decode(&qc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
