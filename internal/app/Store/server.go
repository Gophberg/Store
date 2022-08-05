package Store

import (
	"net/http"
)

var config Config

func Start() error {
	config.NewConfig()
	s := Ad{}
	//http.HandleFunc("/getAd", s.getAd)
	//http.HandleFunc("/getAllAds", s.getAllAds)
	http.HandleFunc("/createAd", s.createAd)
	return http.ListenAndServe(":9000", nil)
}
