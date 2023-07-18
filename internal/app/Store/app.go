package Store

import (
	"fmt"
	"github.com/Gophberg/Store/internal/config"
	"github.com/Gophberg/Store/internal/repository/repo"
	"github.com/Gophberg/Store/internal/usecase"
	"net/http"
)

func Start() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	pg, err := repo.New(&cfg)
	if err != nil {
		return fmt.Errorf("error initializing postgres: %w", err)
	}

	//s := entity.Ad{}

	storeUseCase := usecase.NewUseCase(
		repo.New(pg),
	)

	//var qc *QueryCredentials

	http.HandleFunc("/getAd", s.getAd)
	http.HandleFunc("/getAllAds", s.getAllAds)
	http.HandleFunc("/createAd", s.createAd)

	return http.ListenAndServe(":9000", nil)

}
