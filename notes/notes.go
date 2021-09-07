package notes

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

type Note struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteServer struct {
	Store NoteStore
}

type NoteStore interface {
	GetNote(noteID string) Note
}

func (server NoteServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		server.postNote(rw, r)
	case http.MethodGet:
		server.getNote(rw, r)
	}
}

func (server *NoteServer) getNote(rw http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/users/1/notes/")

	respNote := server.Store.GetNote(noteId)

	rw.Header().Add("Content-Type", "application/json")
	if reflect.DeepEqual(respNote, Note{}) {
		rw.WriteHeader(http.StatusNotFound)
	}

	rwEncoder := json.NewEncoder(rw)
	rwEncoder.Encode(respNote)
}

func (server *NoteServer) postNote(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
}
