package main

import (
	"crud-project/auth"
	"crud-project/handlers"
	"crud-project/middleware"
	"crud-project/models"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	db := models.СonnectDB()
	mux.Handle("/notes", middleware.AuthMiddleware(handlers.NotesHandler(db)))
	mux.Handle("/notes/", middleware.AuthMiddleware(handlers.NotesHandler(db)))
	mux.Handle("/register", auth.RegisterHandler(db))
	mux.Handle("/login", auth.LoginHandler(db))
	mux.Handle("/refresh", handlers.RefreshHandler(db))
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", middleware.LoggingMiddleware(mux.ServeHTTP))

}
