package config

import (
	"os"
	"testing"
)

var configEnvKeys = []string{
	"API_PORT",
	"API_READ_TIMEOUT",
	"API_WRITE_TIMEOUT",
	"API_IDLE_TIMEOUT",
	"API_TIMEOUT",
	"API_API_PORT",
	"API_API_READ_TIMEOUT",
	"API_API_WRITE_TIMEOUT",
	"API_API_IDLE_TIMEOUT",
	"API_API_TIMEOUT",
}

func clearConfigEnv(t *testing.T) {
	t.Helper()
	for _, key := range configEnvKeys {
		key := key
		if prev, ok := os.LookupEnv(key); ok {
			prev := prev
			t.Cleanup(func() {
				_ = os.Setenv(key, prev)
			})
		} else {
			t.Cleanup(func() {
				_ = os.Unsetenv(key)
			})
		}
		_ = os.Unsetenv(key)
	}
}

func TestNewConfig_Defaults(t *testing.T) {
	clearConfigEnv(t)

	cfg, err := NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() error = %v", err)
	}

	if cfg.API.Port != 3000 {
		t.Errorf("API.Port = %d, want 3000", cfg.API.Port)
	}
	if cfg.API.ReadTimeout != 7 {
		t.Errorf("API.ReadTimeout = %d, want 7", cfg.API.ReadTimeout)
	}
	if cfg.API.WriteTimeout != 5 {
		t.Errorf("API.WriteTimeout = %d, want 5", cfg.API.WriteTimeout)
	}
	if cfg.API.IdleTimeout != 5 {
		t.Errorf("API.IdleTimeout = %d, want 5", cfg.API.IdleTimeout)
	}
	if cfg.API.Timeout != 5 {
		t.Errorf("API.Timeout = %d, want 5", cfg.API.Timeout)
	}
}

func TestNewConfig_FromEnv(t *testing.T) {
	t.Setenv("API_PORT", "9090")
	t.Setenv("API_READ_TIMEOUT", "10")
	t.Setenv("API_WRITE_TIMEOUT", "11")
	t.Setenv("API_IDLE_TIMEOUT", "12")
	t.Setenv("API_TIMEOUT", "13")

	cfg, err := NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() error = %v", err)
	}

	if cfg.API.Port != 9090 {
		t.Errorf("API.Port = %d, want 9090", cfg.API.Port)
	}
	if cfg.API.ReadTimeout != 10 {
		t.Errorf("API.ReadTimeout = %d, want 10", cfg.API.ReadTimeout)
	}
	if cfg.API.WriteTimeout != 11 {
		t.Errorf("API.WriteTimeout = %d, want 11", cfg.API.WriteTimeout)
	}
	if cfg.API.IdleTimeout != 12 {
		t.Errorf("API.IdleTimeout = %d, want 12", cfg.API.IdleTimeout)
	}
	if cfg.API.Timeout != 13 {
		t.Errorf("API.Timeout = %d, want 13", cfg.API.Timeout)
	}
}
