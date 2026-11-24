package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"portfolio-backend/data"
)

// ExperienceHandler handles CRUD for experiences
func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store := data.InitStore()

	path := strings.TrimPrefix(r.URL.Path, "/api/experience")
	path = strings.TrimPrefix(path, "/")

	if path != "" {
		switch r.Method {
		case http.MethodPut:
			var update data.Experience
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			updated, err := store.UpdateExperience(path, update)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Experience updated", Data: updated})
			return

		case http.MethodDelete:
			if err := store.DeleteExperience(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Experience deleted"})
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		experiences := store.GetAllExperiences()
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "success",
			Message: "Experience records retrieved successfully",
			Data: map[string]interface{}{
				"count":      len(experiences),
				"experience": experiences,
			},
		})

	case http.MethodPost:
		var newExp data.Experience
		if err := json.NewDecoder(r.Body).Decode(&newExp); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		created, err := store.CreateExperience(newExp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Experience created", Data: created})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// EducationHandler handles CRUD for educations
func EducationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store := data.InitStore()

	path := strings.TrimPrefix(r.URL.Path, "/api/education")
	path = strings.TrimPrefix(path, "/")

	if path != "" {
		switch r.Method {
		case http.MethodPut:
			var update data.Education
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			updated, err := store.UpdateEducation(path, update)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Education updated", Data: updated})
			return

		case http.MethodDelete:
			if err := store.DeleteEducation(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Education deleted"})
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		educations := store.GetAllEducations()
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "success",
			Message: "Education records retrieved successfully",
			Data: map[string]interface{}{
				"count":     len(educations),
				"education": educations,
			},
		})

	case http.MethodPost:
		var newEdu data.Education
		if err := json.NewDecoder(r.Body).Decode(&newEdu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		created, err := store.CreateEducation(newEdu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Education created", Data: created})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// CertificationsHandler handles CRUD for certifications
func CertificationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store := data.InitStore()

	path := strings.TrimPrefix(r.URL.Path, "/api/certifications")
	path = strings.TrimPrefix(path, "/")

	// Sub-route: /api/certifications/next-id
	if path == "next-id" && r.Method == http.MethodGet {
		nextID := store.NextCertificationID()
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"nextId": nextID,
		})
		return
	}

	// Route: /api/certifications/{id}
	if path != "" {
		idInt, err := strconv.Atoi(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: "Invalid certification ID"})
			return
		}

		switch r.Method {
		case http.MethodGet:
			cert := store.GetCertificationByID(idInt)
			if cert == nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: "Certification not found"})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Certification retrieved", Data: cert})
			return

		case http.MethodPut:
			var update data.Certification
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			updated, err := store.UpdateCertification(idInt, update)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Certification updated", Data: updated})
			return

		case http.MethodDelete:
			if err := store.DeleteCertification(idInt); err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
				return
			}
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Certification deleted"})
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	// Route: /api/certifications
	switch r.Method {
	case http.MethodGet:
		categoryQuery := r.URL.Query().Get("category")
		certs := store.GetAllCertifications(categoryQuery)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "success",
			Message: "Certifications retrieved successfully",
			Data: map[string]interface{}{
				"total":          len(store.Certifications),
				"filteredCount":  len(certs),
				"activeCategory": categoryQuery,
				"certifications": certs,
			},
		})

	case http.MethodPost:
		var newCert data.Certification
		if err := json.NewDecoder(r.Body).Decode(&newCert); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		if strings.TrimSpace(newCert.Title) == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: "Title cannot be empty"})
			return
		}
		created, err := store.CreateCertification(newCert)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(APIResponse{Status: "error", Message: err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(APIResponse{Status: "success", Message: "Certification created", Data: created})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ResumeHandler aggregates experience, education, and certifications
func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store := data.InitStore()

	experiences := store.GetAllExperiences()
	educations := store.GetAllEducations()
	certifications := store.GetAllCertifications("")

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(APIResponse{
		Status:  "success",
		Message: "Complete resume data retrieved successfully",
		Data: map[string]interface{}{
			"experience":     experiences,
			"education":      educations,
			"certifications": certifications,
			"summary": map[string]interface{}{
				"totalExperiences":    len(experiences),
				"totalEducations":     len(educations),
				"totalCertifications": len(certifications),
			},
		},
	})
}
