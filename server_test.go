package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETNotes(t *testing.T) {
	t.Run("returns noteid=10 content", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/notes/10", nil)
		response := httptest.NewRecorder()

		NotesServer(response, request)

		got := response.Body.String()
		want := "My random note."

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns noteid=20 content", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/notes/20", nil)
		response := httptest.NewRecorder()

		NotesServer(response, request)

		got := response.Body.String()
		want := "My other note."

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
