package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"portfolio-backend/admin"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// HealthHandler handles root "/" route:
// - If accessed from a web browser (Accept: text/html), serves the v35 Admin Dashboard
// - If accessed as API/cURL (Accept: application/json or non-browser), returns JSON status & sitemap
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// If accessed from a browser, require auth or redirect to login
	if admin.IsBrowserRequest(r) {
		if !admin.IsAuthenticated(r) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		admin.Handler(w, r)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	info := map[string]interface{}{
		"app":       "Portfolio REST API",
		"author":    "Moch Arisandi Jayanto",
		"version":   "1.0.0",
		"adminUI":   "/admin or /dashboard",
		"port":      port,
		"runtime":   "Golang + Cloudflare Workers",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"endpoints": map[string]string{
			"GET /":                   "API status (JSON) or Admin Dashboard (Browser)",
			"GET /admin":              "Modern v35 Admin Dashboard UI",
			"GET /api/health":         "Lightweight health check for FE localhost polling",
			"GET /api/ping":           "Fast ping check",
			"GET /api/projects":       "List all projects (supports ?category=...)",
			"POST /api/projects":      "Create project with auto ID & auto numbering",
			"PUT /api/projects/{id}":  "Update existing project",
			"DELETE /api/projects/{id}": "Delete project",
			"GET /api/certifications": "List all certifications (supports ?category=...)",
			"POST /api/certifications": "Create certification with auto ID",
			"PUT /api/certifications/{id}": "Update certification",
			"DELETE /api/certifications/{id}": "Delete certification",
			"GET /api/resume":         "Aggregated resume data",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(APIResponse{
		Status:  "success",
		Message: "Portfolio API is running smoothly",
		Data:    info,
	})
}

// PingHandler is a lightweight status checker specifically designed for Frontend dev polling
func PingHandler(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "online",
		"service":   "portfolio-backend",
		"port":      port,
		"cors":      "enabled",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
