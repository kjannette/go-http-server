package handler

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/kjannette/go_http_server/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func TestRunServer_ShutdownOnCancel(t *testing.T) {
	e, h := New(zap.NewNop())

	cfg := &config.Config{}
	cfg.API.Port = 0
	cfg.API.ReadTimeout = 2
	cfg.API.WriteTimeout = 2
	cfg.API.IdleTimeout = 2

	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	go func() {
		errCh <- h.RunServer(ctx, cfg, e)
	}()

	healthURL, err := waitForHealthURL(e, 3*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
	resp, err := client.Get(healthURL)
	if err != nil {
		t.Fatalf("GET /health: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if resp.StatusCode != http.StatusOK || string(body) != "OK" {
		t.Fatalf("health check: status=%d body=%q", resp.StatusCode, body)
	}

	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("RunServer() error = %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("RunServer did not exit after context cancel")
	}
}

func waitForHealthURL(e *echo.Echo, timeout time.Duration) (string, error) {
	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if e.Listener != nil {
			_, port, err := net.SplitHostPort(e.Listener.Addr().String())
			if err != nil {
				return "", fmt.Errorf("split listener addr: %w", err)
			}
			url := fmt.Sprintf("http://127.0.0.1:%s/health", port)

			resp, err := client.Get(url)
			if err == nil {
				_, _ = io.Copy(io.Discard, resp.Body)
				_ = resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					return url, nil
				}
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	return "", fmt.Errorf("server not ready within %s", timeout)
}
