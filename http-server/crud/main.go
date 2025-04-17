package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loggingMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
func usersHandler(users *map[int]User, nextID *int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/users" {
				var result []User
				for _, user := range *users {
					result = append(result, user)
				}
				json.NewEncoder(w).Encode(result)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/users/") {
				idStr := strings.TrimPrefix(r.URL.Path, "/users/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid user ID", http.StatusBadRequest)
					return
				}
				user, ok := (*users)[id]
				if !ok {
					http.Error(w, "User not found", http.StatusNotFound)
					return
				}
				json.NewEncoder(w).Encode(user)
				return
			}
			http.NotFound(w, r)
		case http.MethodPost:
			var newUser User
			if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				return
			}
			if newUser.Age <= 0 || newUser.Name == "" {
				http.Error(w, "Invalid user data", http.StatusBadRequest)
				return
			}
			newUser.ID = *nextID
			(*users)[*nextID] = newUser
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newUser)
			return
		case http.MethodPut:
			idStr := strings.TrimPrefix(r.URL.Path, "/users/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			if _, ok := (*users)[id]; ok {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			var updateUser User
			if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				return
			}
			if updateUser.Age <= 0 || updateUser.Name == "" {
				http.Error(w, "Invalid user data", http.StatusBadRequest)
				return
			}
			updateUser.ID = id
			(*users)[id] = updateUser
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(updateUser)
		case http.MethodDelete:
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
}

func main() {
	mux := http.NewServeMux()
	var users = make(map[int]User)
	var nextID = 1
	users[nextID] = User{ID: nextID, Name: "John Doe", Age: 25}
	nextID++
	users[nextID] = User{ID: nextID, Name: "Alan Roe", Age: 26}
	nextID++
	mux.Handle("/users", usersHandler(&users, &nextID))
	mux.Handle("/users/", usersHandler(&users, &nextID))
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", loggingMiddleware(mux.ServeHTTP))
}
