package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"portfolio-backend/data"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store := data.InitStore()

	// Extract subpath: e.g. /api/projects or /api/projects/{id}
	path := strings.TrimPrefix(r.URL.Path, "/api/projects")
	path = strings.TrimPrefix(path, "/")

	// Sub-route: /api/projects/next-sequence
	if path == "next-sequence" && r.Method == http.MethodGet {
		num, seq := store.NextProjectSequence()
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     "success",
			"nextNumber": num,
			"sequence":   seq,
		})
		return
	}

	// Route: /api/projects/{id}
	if path != "" {
		switch r.Method {
		case http.MethodGet:
			project := store.GetProjectByID(path)
			if project == nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{
					Status:  "error",
					Message: "Project with ID '" + path + "' not found",
				})
				return
			}
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "success",
				Message: "Project detail retrieved",
				Data:    project,
			})
			return

		case http.MethodPut:
			var update data.Project
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(APIResponse{
					Status:  "error",
					Message: "Invalid request payload: " + err.Error(),
				})
				return
			}
			updatedProj, err := store.UpdateProject(path, update)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{
					Status:  "error",
					Message: err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "success",
				Message: "Project updated successfully",
				Data:    updatedProj,
			})
			return

		case http.MethodDelete:
			if err := store.DeleteProject(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(APIResponse{
					Status:  "error",
					Message: err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "success",
				Message: "Project deleted successfully",
			})
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	// Route: /api/projects
	switch r.Method {
	case http.MethodGet:
		categoryQuery := r.URL.Query().Get("category")
		projects := store.GetAllProjects(categoryQuery)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "success",
			Message: "Projects retrieved successfully",
			Data: map[string]interface{}{
				"count":    len(projects),
				"projects": projects,
			},
		})

	case http.MethodPost:
		var newProj data.Project
		if err := json.NewDecoder(r.Body).Decode(&newProj); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "error",
				Message: "Invalid request payload: " + err.Error(),
			})
			return
		}

		if strings.TrimSpace(newProj.Title) == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "error",
				Message: "Title cannot be empty",
			})
			return
		}

		createdProj, err := store.CreateProject(newProj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "error",
				Message: "Failed to create project: " + err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "success",
			Message: "Project created successfully",
			Data:    createdProj,
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
