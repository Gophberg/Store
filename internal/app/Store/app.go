package Store

import (
	"github.com/Gophberg/Store/internal/config"
	"net/http"
)

func Start() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	s := Ad{}

	//var qc *QueryCredentials

	http.HandleFunc("/getAd", s.getAd)
	http.HandleFunc("/getAllAds", s.getAllAds)
	http.HandleFunc("/createAd", s.createAd)

	return http.ListenAndServe(":9000", nil)

}
