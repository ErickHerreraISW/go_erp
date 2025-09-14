package middleware

import (
	"context"
	"net/http"
)

// RequestIDToContext Guarda el request_id en el contexto para logs
func RequestIDToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		ctx := context.WithValue(r.Context(), "request_id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
