package http

import (
	"net/http"
	"time"

	ih "github.com/ErickHerreraISW/go_erp/internal/http/middleware"
	"github.com/ErickHerreraISW/go_erp/internal/products"
	"github.com/ErickHerreraISW/go_erp/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type RouterDeps struct {
	UserHandler    *users.Handler
	ProductHandler *products.Handler
	JWTAuth        *jwtauth.JWTAuth
}

func NewRouter(d RouterDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(ih.RequestIDToContext)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	r.Route("/v1", func(r chi.Router) {
		// Auth
		r.Post("/auth/login", d.UserHandler.Login)

		// Users
		r.Route("/users", func(r chi.Router) {
			r.Post("/", d.UserHandler.Create)
			r.Get("/", d.UserHandler.List)
			r.Get("/{id}", d.UserHandler.GetByID)
			r.With(jwtauth.Verifier(d.JWTAuth), jwtauth.Authenticator(d.JWTAuth)).Put("/{id}", d.UserHandler.Update)
			r.With(jwtauth.Verifier(d.JWTAuth), jwtauth.Authenticator(d.JWTAuth)).Delete("/{id}", d.UserHandler.Delete)
		})

		// Products
		r.Route("/products", func(r chi.Router) {
			r.With(jwtauth.Verifier(d.JWTAuth), jwtauth.Authenticator(d.JWTAuth)).Post("/", d.ProductHandler.Create)
			r.Get("/", d.ProductHandler.List)
			r.Get("/{id}", d.ProductHandler.GetByID)
			r.With(jwtauth.Verifier(d.JWTAuth), jwtauth.Authenticator(d.JWTAuth)).Put("/{id}", d.ProductHandler.Update)
			r.With(jwtauth.Verifier(d.JWTAuth), jwtauth.Authenticator(d.JWTAuth)).Delete("/{id}", d.ProductHandler.Delete)
		})
	})

	return r
}
