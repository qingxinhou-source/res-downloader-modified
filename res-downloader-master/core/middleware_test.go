package core

import (
	"net/http/httptest"
	"testing"
)

func TestIsAllowedAPIOrigin(t *testing.T) {
	tests := []struct {
		name    string
		origin  string
		allowed bool
	}{
		{name: "empty origin", allowed: true},
		{name: "wails webview", origin: "http://wails.localhost", allowed: true},
		{name: "localhost dev server", origin: "http://localhost:5173", allowed: true},
		{name: "loopback api host", origin: "http://127.0.0.1:8899", allowed: true},
		{name: "ipv6 loopback", origin: "http://[::1]:8899", allowed: true},
		{name: "external website", origin: "https://example.com", allowed: false},
		{name: "spoofed suffix", origin: "https://127.0.0.1.example.com", allowed: false},
		{name: "invalid origin", origin: "://bad-origin", allowed: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/app-info", nil)
			if tt.origin != "" {
				req.Header.Set("Origin", tt.origin)
			}

			if got := isAllowedAPIOrigin(req); got != tt.allowed {
				t.Fatalf("isAllowedAPIOrigin() = %v, want %v", got, tt.allowed)
			}
		})
	}
}
