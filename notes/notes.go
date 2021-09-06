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
}

func (server NoteServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/users/1/notes/")
	respNote := Note{
		ID:      "1001",
		UserID:  "1",
		Title:   "Awesome note",
		Content: "My awesome note.",
	}

	if noteId == "2002" {
		respNote = Note{
			ID:      "2002",
			UserID:  "1",
			Title:   "Other note",
			Content: "My other note.",
		}
	}

	rwEncoder := json.NewEncoder(rw)
	rwEncoder.Encode(respNote)
}
