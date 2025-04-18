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
	resp, err := http.Get("http://localhost:8080/notes")
	if err != nil {
		log.Fatal(err)
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
	resp, err := http.Post("http://localhost:8080/notes", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Request error", err)
		return
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

}

func deleteNote(note Note) {

}

func main() {
	note := Note{
		Title:   "From client",
		Content: "Created via POST",
	}
	getNotes()
	createNote(note)
	getNotes()
}
