package main

import (
	"fmt"
	"net/http"
	"strings"
)

func NotesServer(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")

	if noteId == "10" {
		fmt.Fprintf(w, "My random note.")
		return
	}

	if noteId == "20" {
		fmt.Fprintf(w, "My other note.")
		return
	}
}
