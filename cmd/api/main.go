package main

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	"github.com/youruser/myapp/internal/config"
	"github.com/youruser/myapp/internal/database"
	apphttp "github.com/youruser/myapp/internal/http"
	"github.com/youruser/myapp/internal/logger"
	"github.com/youruser/myapp/internal/products"
	"github.com/youruser/myapp/internal/users"
)

func main() {
	cfg := config.Load()
	logger.Setup(cfg.AppEnv)

	db, err := database.New(cfg.DBURL)
	if err != nil {
		log.Fatal().Err(err).Msg("db connection failed")
	}

	// Auto-migraciones mínimas
	if err := db.AutoMigrate(&users.User{}, &products.Product{}); err != nil {
		log.Fatal().Err(err).Msg("migrations failed")
	}

	// DI manual
	usrRepo := users.NewRepository(db)
	usrSvc := users.NewService(usrRepo)

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
