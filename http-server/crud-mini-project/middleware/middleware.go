package middleware

import (
	"context"
	"crud-project/auth"
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
		claims := &auth.Claims{}
		token, err := jwt.Parse(tokerStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return auth.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userID := claims.UserID
		ctx := context.WithValue(r.Context(), UserIdKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
