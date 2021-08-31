package main

import (
	"fmt"
	"net/http"
	"strings"
)

func NotesServer(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")
	fmt.Fprintf(w, GetNote(noteId))
}

func GetNote(noteId string) string {
	if noteId == "10" {
		return "My random note."
	}

	if noteId == "20" {
		return "My other note."
	}

	return ""
}
