package users

import (
	"github.com/ErickHerreraISW/go_erp/internal/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterRoutes(r chi.Router, h *Handler, jwt *jwtauth.JWTAuth) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.List)
		r.Get("/{id}", h.GetByID)

		secure := func(r chi.Router) {
			r.Use(jwtauth.Verifier(jwt))
			r.Use(jwtauth.Authenticator(jwt))
			r.Use(middleware.AuthenticatedUser[User](h.Svc))
		}

		r.Group(func(r chi.Router) {
			secure(r)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}
