package main

import (
	"log"
	"net/http"

	"github.com/mehrankamal/noter/notes"
)

type InMemoryNoteStore struct{}

func (store *InMemoryNoteStore) GetNote(noteID string) notes.Note {
	return notes.Note{
		ID:      noteID,
		UserID:  "1",
		Title:   "Awesome Title",
		Content: "My awesome note",
	}
}

func main() {
	server := &notes.NoteServer{&InMemoryNoteStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
