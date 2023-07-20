package Store

import (
	"fmt"
	"github.com/Gophberg/Store/internal/config"
	v1 "github.com/Gophberg/Store/internal/handler/http/api/v1"
	"github.com/Gophberg/Store/internal/repository/repo"
	"github.com/Gophberg/Store/internal/usecase"
	"github.com/Gophberg/Store/pkg/postgres"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	router *mux.Router
}

func Start() error {

	var a App

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

	// HandlerUseCase
	storeUseCase := usecase.NewUseCase(
		repo.New(pg),
	)

	handler := v1.NewHandler(storeUseCase)
	a.router = mux.NewRouter()
	a.router.HandleFunc("/createAd", handler.CreateAd).Methods("POST")
	a.router.HandleFunc("/getAllAds", handler.GetAllAds).Methods("GET")
	a.router.HandleFunc("/getAd", handler.GetAd).Methods("GET")

	return http.ListenAndServe(":9000", a.router)
}
