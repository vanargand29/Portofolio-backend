package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"portfolio-backend/data"
)

// UploadImageHandler handles POST /api/upload for uploading images to Cloudflare KV or local disk
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Max 10MB upload limit
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "error",
			Message: "Ukuran file terlalu besar atau format multipart tidak valid (maksimal 10MB)",
		})
		return
	}

	// Try "file" field first, then fallback to "image"
	file, header, err := r.FormFile("file")
	if err != nil {
		file, header, err = r.FormFile("image")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(APIResponse{
				Status:  "error",
				Message: "Field file gambar tidak ditemukan ('file' atau 'image')",
			})
			return
		}
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "error",
			Message: "Gagal membaca isi file gambar",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext == "" {
		ext = ".png"
	}
	baseName := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
	cleanBase := data.Slugify(baseName)
	if cleanBase == "" {
		cleanBase = "upload"
	}

	// Unique filename with timestamp
	fileName := fmt.Sprintf("%d-%s%s", time.Now().Unix(), cleanBase, ext)

	// Save to Cloudflare KV / local disk under "img/<filename>"
	err = data.SaveToKV("img/"+fileName, fileBytes)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(APIResponse{
			Status:  "error",
			Message: "Gagal menyimpan gambar ke Cloudflare Storage: " + err.Error(),
		})
		return
	}

	imageURL := "/api/images/" + fileName

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(APIResponse{
		Status:  "success",
		Message: "Gambar berhasil diunggah dan disimpan ke Cloudflare",
		Data: map[string]string{
			"url":      imageURL,
			"filename": fileName,
		},
	})
}

// ServeImageHandler serves images from Cloudflare KV (or local disk fallback)
func ServeImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Strip prefix to extract filename
	path := r.URL.Path
	fileName := strings.TrimPrefix(path, "/api/images/")
	fileName = strings.TrimPrefix(fileName, "/img/")
	fileName = strings.TrimSpace(fileName)

	if unescaped, err := url.PathUnescape(fileName); err == nil && unescaped != "" {
		fileName = unescaped
	}

	if fileName == "" {
		http.NotFound(w, r)
		return
	}

	// 1. Try loading from Cloudflare KV
	imgBytes, err := data.LoadFromKV("img/" + fileName)
	if err == nil && len(imgBytes) > 0 {
		serveBytes(w, fileName, imgBytes)
		return
	}

	// 2. Fallback to local files if running locally
	localPaths := []string{
		filepath.Join("data", "img", fileName),
		filepath.Join("data", fileName),
		filepath.Join("public", "img", fileName),
	}
	for _, lp := range localPaths {
		if b, err := os.ReadFile(lp); err == nil && len(b) > 0 {
			serveBytes(w, fileName, b)
			return
		}
	}

	http.NotFound(w, r)
}

func serveBytes(w http.ResponseWriter, fileName string, b []byte) {
	ext := strings.ToLower(filepath.Ext(fileName))
	contentType := "application/octet-stream"
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".svg":
		contentType = "image/svg+xml"
	case ".webp":
		contentType = "image/webp"
	case ".gif":
		contentType = "image/gif"
	case ".ico":
		contentType = "image/x-icon"
	default:
		contentType = http.DetectContentType(b)
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
