package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	helloHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "Hello!") {
		t.Errorf("Unknown Body: %s", body)
	}
}
