package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type UserResponse struct {
	Message string `json:"message"`
}
type StatusResponse struct {
	Status string `json:"status"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := StatusResponse{
		Status: "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var message MessageResponse
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
func loggingMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" || user.Age <= 0 {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(user)
	response := MessageResponse{
		Message: fmt.Sprintf("User %s added!", user.Name),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", http.HandlerFunc(userHandler))
	mux.Handle("/ping", http.HandlerFunc(pingHandler))
	mux.Handle("/greet", http.HandlerFunc(greetHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/echo", http.HandlerFunc(echoHandler))
	mux.Handle("/", http.HandlerFunc(helloHandler))
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", loggingMiddleware(mux.ServeHTTP))
}
