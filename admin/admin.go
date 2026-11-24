package admin

import (
	_ "embed"
	"net/http"
	"strings"
)

//go:embed index.html
var DashboardHTML string

// Handler serves the modern v35 Admin Dashboard UI
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(DashboardHTML))
}

// IsBrowserRequest checks if the incoming request is from a web browser requesting HTML
func IsBrowserRequest(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "text/html")
}
