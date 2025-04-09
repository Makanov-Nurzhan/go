package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// создаём безопасный тип ключа
type ctxKey string

const userIDKey ctxKey = "userID"

// middleware добавляет userID в контекст запроса
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// допустим, пользователь авторизован как ID = 42
		ctx := context.WithValue(r.Context(), userIDKey, 42)

		// передаём обновлённый контекст дальше
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// основной обработчик, извлекающий userID из контекста
func handler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey)

	if userID != nil {
		fmt.Fprintf(w, "👤 Hello, user #%v!\n", userID)
		log.Println("✅ Обработка запроса для userID:", userID)
	} else {
		http.Error(w, "🚫 No user ID", http.StatusUnauthorized)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", authMiddleware(http.HandlerFunc(handler)))

	fmt.Println("🚀 Слушаю на :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
