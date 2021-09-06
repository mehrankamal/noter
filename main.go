package main

import (
	"log"
	"net/http"

	"github.com/mehrankamal/noter/notes"
)

func main() {
	server := &notes.NoteServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
