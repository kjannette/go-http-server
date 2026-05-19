package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestHealthCheck(t *testing.T) {
	e, _ := New(zap.NewNop())

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(body) != "OK" {
		t.Errorf("body = %q, want %q", body, "OK")
	}

	if ct := rec.Header().Get("Content-Type"); ct != "text/plain; charset=UTF-8" {
		t.Errorf("Content-Type = %q, want text/plain; charset=UTF-8", ct)
	}
}
