package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
)

func NewAuthorization(cfg config.Config) func(next http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authorizationHeader, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			tokenString := authorizationHeader[7:]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Secret), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user-id", claims["id"].(string))))
		})
	}
}
