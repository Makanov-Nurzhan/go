package handlers

import (
	"crud-project/models"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func NotesHandler(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/notes" {
				var notes []models.Note
				result := db.Find(&notes)
				if result.Error != nil {
					http.Error(w, "Database error", http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(notes)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/notes/") {
				idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}
				var note models.Note
				result := db.First(&note, id)
				if result.Error != nil {
					http.Error(w, "Note not found", http.StatusNotFound)
					return
				}
				json.NewEncoder(w).Encode(note)
				return
			}
			http.NotFound(w, r)
		case http.MethodPost:
			var newNote models.Note
			if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			if newNote.Title == "" || newNote.Content == "" {
				http.Error(w, "Invalid data of note", http.StatusBadRequest)
				return
			}
			result := db.Create(&newNote)
			if result.Error != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newNote)
			return
		case http.MethodPut:
			idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			var existingNote models.Note
			if err := db.First(&existingNote, id).Error; err != nil {
				http.Error(w, "Note not found", http.StatusNotFound)
				return
			}

			var updateNote models.Note
			if err := json.NewDecoder(r.Body).Decode(&updateNote); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			if updateNote.Title == "" || updateNote.Content == "" {
				http.Error(w, "Invalid data of note", http.StatusBadRequest)
				return
			}
			existingNote.Title = updateNote.Title
			existingNote.Content = updateNote.Content
			if err := db.Save(&existingNote).Error; err != nil {
				http.Error(w, "Database update error", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(existingNote)
			return
		case http.MethodDelete:
			if !strings.HasPrefix(r.URL.Path, "/notes/") {
				http.Error(w, "Invalid note id", http.StatusBadRequest)
				return
			}
			idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			var note models.Note
			if err := db.First(&note, id).Error; err != nil {
				http.Error(w, "Note not found", http.StatusNotFound)
				return
			}
			db.Delete(&note)
			w.WriteHeader(http.StatusNoContent)
			return
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})
}
