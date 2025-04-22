package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

var jwtSecret = []byte("secretkey")

type contextKey string

const UserIdKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokerStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokerStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}
		userID := uint(claims["user_id"].(float64))
		ctx := context.WithValue(r.Context(), UserIdKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
