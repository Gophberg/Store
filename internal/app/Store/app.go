package Store

import (
	"fmt"
	"github.com/Gophberg/Store/internal/config"
	"github.com/Gophberg/Store/internal/repository/repo"
	"github.com/Gophberg/Store/internal/usecase"
	"github.com/Gophberg/Store/pkg/postgres"
)

func Start() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Dbusername, cfg.Dbpassword, cfg.Dbhost, cfg.Dockerdbport, cfg.Dbname)

	// Repository
	pg, err := postgres.New(url, postgres.MaxPoolSize(cfg.MaxPoolSize), postgres.ConnAttempts(cfg.ConnAttempts), postgres.ConnTimeout(cfg.ConnTimeout))
	if err != nil {
		return fmt.Errorf("error initializing postgres: %w", err)
	}
	defer pg.Close()

	// UseCase
	storeUseCase := usecase.NewUseCase(
		repo.New(pg),
	)

	//// Handler
	//http.HandleFunc("/getAd", s.getAd)
	//http.HandleFunc("/getAllAds", s.getAllAds)
	//http.HandleFunc("/createAd", s.createAd)
	//
	//return http.ListenAndServe(":9000", nil)

}
