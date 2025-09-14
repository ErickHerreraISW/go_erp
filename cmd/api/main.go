package main

import (
	"net/http"

	"github.com/ErickHerreraISW/go_erp/internal/config"
	"github.com/ErickHerreraISW/go_erp/internal/database"
	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstance"
	"github.com/ErickHerreraISW/go_erp/internal/feature/products"
	"github.com/ErickHerreraISW/go_erp/internal/feature/users"
	apphttp "github.com/ErickHerreraISW/go_erp/internal/http"
	"github.com/ErickHerreraISW/go_erp/internal/logger"
	"github.com/ErickHerreraISW/go_erp/internal/platform/migrations"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load()
	logger.Setup("Cargando .env")
	logger.Setup(cfg.AppEnv)

	db, err := database.New(cfg.DBURL)
	if err != nil {
		log.Fatal().Err(err).Msg("db connection failed")
	}

	// Auto-migraciones mínimas
	if err := migrations.Run(db); err != nil {
		log.Fatal().Err(err).Msg("migrations failed")
	}

	// DI manual
	erpInsRepo := erpinstance.NewRepository(db)
	//erpInsSvc := erpinstance.NewService(erpInsRepo)

	usrRepo := users.NewRepository(db)
	usrSvc := users.NewService(usrRepo, erpInsRepo)

	prdRepo := products.NewRepository(db)
	prdSvc := products.NewService(prdRepo)

	jwtAuth := jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	usrHandler := users.NewHandler(usrSvc, jwtAuth)
	prdHandler := products.NewHandler(prdSvc)

	r := apphttp.NewRouter(apphttp.RouterDeps{
		UserHandler:    usrHandler,
		ProductHandler: prdHandler,
		JWTAuth:        jwtAuth,
	})

	addr := ":" + cfg.HTTPPort
	log.Info().Msgf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("server failed")
	}
}
