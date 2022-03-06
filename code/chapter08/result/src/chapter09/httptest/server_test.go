package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	http.DefaultServeMux.HandleFunc("/topic/", handleRequest)

	w = httptest.NewRecorder()

	os.Exit(m.Run())
}

func TestHandlePost(t *testing.T) {
	reader := strings.NewReader(`{"title":"The Go Standard Library","content":"It contains many packages."}`)
	r, _ := http.NewRequest(http.MethodPost, "/topic/", reader)

	http.DefaultServeMux.ServeHTTP(w, r)

	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", result.StatusCode)
	}
}

func TestHandleGet(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/topic/1", nil)

	http.DefaultServeMux.ServeHTTP(w, r)

	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", result.StatusCode)
	}

	topic := new(Topic)
	json.Unmarshal(w.Body.Bytes(), topic)
	if topic.Id != 1 {
		t.Errorf("Cannot get topic")
	}
}
