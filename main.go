package main

import (
	"log"
	"net/http"
)

type InMemoryNoteStore struct{}

func (inMem *InMemoryNoteStore) GetNote(noteId string) string {
	return "My awesome note"
}

func main() {
	server := &NoteServer{&InMemoryNoteStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
