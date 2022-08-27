package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := NewMux()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("expect status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	expect := `{"status": "ok"}`
	if string(got) != expect {
		t.Errorf("expect %q, but got %q", expect, got)
	}
}
