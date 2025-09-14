package http

import (
	"net/http"
	"time"

	feaprod "github.com/ErickHerreraISW/go_erp/internal/feature/products"
	feausr "github.com/ErickHerreraISW/go_erp/internal/feature/users"
	ih "github.com/ErickHerreraISW/go_erp/internal/http/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type RouterDeps struct {
	UserHandler    *feausr.Handler
	ProductHandler *feaprod.Handler
	JWTAuth        *jwtauth.JWTAuth
}

func NewRouter(d RouterDeps) http.Handler {
	r := chi.NewRouter()

	// Middlewares globales
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(ih.RequestIDToContext)

	// Health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("ok")) })

	// Versión
	r.Route("/v1", func(api chi.Router) {
		// Auth
		api.Post("/auth/login", d.UserHandler.Login)

		// Features
		feausr.RegisterRoutes(api, d.UserHandler, d.JWTAuth)
		feaprod.RegisterRoutes(api, d.ProductHandler, d.JWTAuth)
	})

	return r
}
