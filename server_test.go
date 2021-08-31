package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETNotes(t *testing.T) {
	t.Run("returns noteid=10 content", func(t *testing.T) {
		request := newGetNoteRequest("10")
		response := httptest.NewRecorder()

		NotesServer(response, request)

		assertResponseBody(t, response.Body.String(), "My random note.")
	})
	t.Run("returns noteid=20 content", func(t *testing.T) {

		request := newGetNoteRequest("20")
		response := httptest.NewRecorder()

		NotesServer(response, request)

		assertResponseBody(t, response.Body.String(), "My other note.")
	})
}

func newGetNoteRequest(noteId string) *http.Request {
	reqUrl := fmt.Sprintf("/notes/%s", noteId)
	req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong response body, got %q want %q", got, want)
	}
}
