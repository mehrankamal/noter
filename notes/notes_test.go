package notes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubNoteStore struct {
	notes map[string]Note
}

func (s *StubNoteStore) GetNote(noteID string) Note {
	note := s.notes[noteID]
	return note
}

func TestGETNotes(t *testing.T) {
	store := StubNoteStore{
		map[string]Note{
			"1001": {
				ID:      "1001",
				UserID:  "1",
				Title:   "Awesome note",
				Content: "My awesome note.",
			},
			"2002": {
				ID:      "2002",
				UserID:  "1",
				Title:   "Other note",
				Content: "My other note.",
			}}}

	server := NoteServer{&store}

	t.Run("get note id `1001` for user 1", func(t *testing.T) {
		request := newGetNoteRequest("1", "1001")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		expectedNote := Note{
			ID:      "1001",
			UserID:  "1",
			Title:   "Awesome note",
			Content: "My awesome note.",
		}
		assertResponseBody(t, response.Body, expectedNote)
	})
	t.Run("get note id 2002 for user 1", func(t *testing.T) {
		request := newGetNoteRequest("1", "2002")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		expectedNote := Note{
			ID:      "2002",
			UserID:  "1",
			Title:   "Other note",
			Content: "My other note.",
		}
		assertResponseBody(t, response.Body, expectedNote)
	})
}

func newGetNoteRequest(userId, noteId string) *http.Request {
	reqUrl := fmt.Sprintf("/users/%s/notes/%s", userId, noteId)
	req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
	return req
}

func assertResponseBody(t testing.TB, respBody *bytes.Buffer, expectedNote Note) {
	var gotNote Note

	respDecoder := json.NewDecoder(respBody)
	respDecoder.Decode(&gotNote)

	if !reflect.DeepEqual(gotNote, expectedNote) {
		t.Errorf("Expected %v, got %v", expectedNote, gotNote)
	}
}
