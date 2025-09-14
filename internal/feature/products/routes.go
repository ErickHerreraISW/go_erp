package products

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterRoutes(r chi.Router, h *Handler, jwt *jwtauth.JWTAuth) {
	r.Route("/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(jwt))
			r.Use(jwtauth.Authenticator(jwt))
			r.Post("/", h.Create)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})

		r.Get("/", h.List)
		r.Get("/{id}", h.GetByID)
	})
}
