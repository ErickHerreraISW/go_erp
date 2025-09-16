package middleware

import (
	"context"
	"net/http"

	"github.com/ErickHerreraISW/go_erp/internal/core/contracts"
	"github.com/go-chi/jwtauth/v5"
)

// RequestIDToContext Guarda el request_id en el contexto para logs
func RequestIDToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		ctx := context.WithValue(r.Context(), "request_id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ctxKeyUser struct{}

func AuthenticatedUser[T any](up contracts.UserProvider[T]) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			sub, ok := claims["sub"].(float64)
			if !ok {
				http.Error(w, "invalid token", 401)
				return
			}

			u, err := up.Get(uint(sub))
			if err != nil || u == nil {
				http.Error(w, "user not found", 401)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeyUser{}, u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserFromContext[T any](ctx context.Context) (*T, bool) {
	u, ok := ctx.Value(ctxKeyUser{}).(*T)
	return u, ok
}
