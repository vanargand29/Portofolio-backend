package main

import (
	"log"
	"net/http"
	"os"

	"github.com/syumai/workers-go"
	"portfolio-backend/admin"
	"portfolio-backend/data"
	"portfolio-backend/handlers"
	"portfolio-backend/middleware"
)

func main() {
	// Initialize data store with default seed or persisted store.json
	data.InitStore()

	// Set default port to 8080 for local development if not specified
	if os.Getenv("PORT") == "" {
		_ = os.Setenv("PORT", "8080")
	}

	mux := http.NewServeMux()

	// Register Admin Dashboard endpoints (served directly from portofolio-backend/admin)
	mux.HandleFunc("/", handlers.HealthHandler)
	mux.HandleFunc("/admin/login", admin.LoginHandler)
	mux.HandleFunc("/admin/logout", admin.LogoutHandler)
	mux.HandleFunc("/admin", admin.RequireAuth(admin.Handler))
	mux.HandleFunc("/admin/", admin.RequireAuth(admin.Handler))
	mux.HandleFunc("/dashboard", admin.RequireAuth(admin.Handler))
	mux.HandleFunc("/dashboard/", admin.RequireAuth(admin.Handler))

	// Register API endpoints (with /api/ prefix and direct root prefix for custom domain api.mocharisandijayanto.my.id)
	mux.HandleFunc("/api/health", handlers.PingHandler)
	mux.HandleFunc("/api/ping", handlers.PingHandler)
	mux.HandleFunc("/api/projects", handlers.ProjectsHandler)
	mux.HandleFunc("/api/projects/", handlers.ProjectsHandler)
	mux.HandleFunc("/api/experience", handlers.ExperienceHandler)
	mux.HandleFunc("/api/experience/", handlers.ExperienceHandler)
	mux.HandleFunc("/api/education", handlers.EducationHandler)
	mux.HandleFunc("/api/education/", handlers.EducationHandler)
	mux.HandleFunc("/api/certifications", handlers.CertificationsHandler)
	mux.HandleFunc("/api/certifications/", handlers.CertificationsHandler)
	mux.HandleFunc("/api/resume", handlers.ResumeHandler)
	mux.HandleFunc("/api/upload", handlers.UploadImageHandler)
	mux.HandleFunc("/api/images/", handlers.ServeImageHandler)
	mux.HandleFunc("/img/", handlers.ServeImageHandler)

	// Direct endpoint aliases (e.g. https://api.mocharisandijayanto.my.id/projects)
	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/projects", handlers.ProjectsHandler)
	mux.HandleFunc("/projects/", handlers.ProjectsHandler)
	mux.HandleFunc("/experience", handlers.ExperienceHandler)
	mux.HandleFunc("/experience/", handlers.ExperienceHandler)
	mux.HandleFunc("/education", handlers.EducationHandler)
	mux.HandleFunc("/education/", handlers.EducationHandler)
	mux.HandleFunc("/certifications", handlers.CertificationsHandler)
	mux.HandleFunc("/certifications/", handlers.CertificationsHandler)
	mux.HandleFunc("/resume", handlers.ResumeHandler)
	mux.HandleFunc("/upload", handlers.UploadImageHandler)

	// Wrap router with CORS middleware
	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	log.Println("🚀 Portfolio Backend API & Admin Dashboard is ready")
	log.Printf("⚡ Running on port %s / Cloudflare Workers environment\n", port)
	log.Printf("👉 Admin Dashboard: http://localhost:%s/admin\n", port)

	// workers.Serve handles both native Go execution and Cloudflare Workers runtime
	workers.Serve(handler)
}
