package notes

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Note struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteServer struct {
	store NoteStore
}

type NoteStore interface {
	GetNote(noteID string) Note
}

func (server NoteServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/users/1/notes/")

	respNote := server.store.GetNote(noteId)

	rwEncoder := json.NewEncoder(rw)
	rwEncoder.Encode(respNote)
}
