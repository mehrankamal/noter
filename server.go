package main

import (
	"fmt"
	"net/http"
	"strings"
)

type NoteStore interface {
	GetNote(noteId string) string
}

type NoteServer struct {
	store NoteStore
}

func (n *NoteServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")

	note := n.store.GetNote(noteId)

	if note == "" {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, note)
}
