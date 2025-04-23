package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getNotes() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/notes", nil)
	if err != nil {
		log.Fatal("Request creation failed: ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}
	defer resp.Body.Close()
	var notes []Note
	if err := json.NewDecoder(resp.Body).Decode(&notes); err != nil {
		log.Fatal("Decode error", err)
	}
	for _, note := range notes {
		fmt.Printf("Note ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
	}

}

func createNote(note Note) {
	jsonData, err := json.Marshal(note)
	if err != nil {
		log.Fatal("Encode error", err)
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/notes", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Request creation failed: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}
	defer resp.Body.Close()
	var createdNote Note
	err = json.NewDecoder(resp.Body).Decode(&createdNote)
	if err != nil {
		log.Fatal("Decode error", err)
		return
	}
	fmt.Printf("Note ID: %d, Title: %s, Content: %s\n", createdNote.ID, createdNote.Title, createdNote.Content)
}

func updateNote(note Note) {
	jsonData, err := json.Marshal(note)
	if err != nil {
		log.Fatal("Encode error", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/notes/%d", note.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Request error", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request error", err)
	}
	defer resp.Body.Close()
	var updatedNote Note
	if err := json.NewDecoder(resp.Body).Decode(&updatedNote); err != nil {
		log.Fatal("Decode error", err)
	}
	if resp.StatusCode != http.StatusAccepted {
		fmt.Printf("Update failed: %s\n", resp.Status)
	}

}

func deleteNote(id int) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/notes/%d", id), nil)
	if err != nil {
		log.Fatal("Request error", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request error", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Note Deleted")
	} else {
		fmt.Printf("Failed to delete note %d, status %s", id, resp.Status)
	}
}

func main() {
	note := Note{
		Title:   "From client",
		Content: "Created via POST",
	}
	updatedNote := Note{ID: 3, Title: "Updated", Content: "Updated from client"}
	getNotes()
	fmt.Println("Get Notes")
	createNote(note)
	fmt.Println("Create Note")
	getNotes()
	fmt.Println("Get after creating Note")
	updateNote(updatedNote)
	fmt.Println("Update Note")
	getNotes()
	fmt.Println("Get after update Note")
	deleteNote(3)
	fmt.Println("Delete Note")
	getNotes()
	fmt.Println("Get after delete Note")
}
