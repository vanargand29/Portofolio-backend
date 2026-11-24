package data

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

// Store holds all portfolio data in memory with thread safety and Cloudflare KV persistence
type Store struct {
	mu             sync.RWMutex
	Projects       []Project       `json:"projects"`
	Experiences    []Experience    `json:"experiences"`
	Educations     []Education     `json:"educations"`
	Certifications []Certification `json:"certifications"`
	filePath       string
}

var (
	GlobalStore *Store
	once        sync.Once
)

const (
	defaultStorePath = "data/store.json"
	kvStoreKey       = "store_data"
)

// InitStore initializes the global data store with default seeds or from store.json
func InitStore() *Store {
	once.Do(func() {
		GlobalStore = &Store{
			filePath: defaultStorePath,
		}
		// 1. Try to load from store.json if exists
		if err := GlobalStore.loadFromFile(); err != nil {
			// Initialize with default seeds
			GlobalStore.Projects = Projects
			GlobalStore.Experiences = Experiences
			GlobalStore.Educations = Educations
			GlobalStore.Certifications = Certifications
		}
	})
	return GlobalStore
}

func (s *Store) loadFromFile() error {
	if s.filePath == "" {
		return fmt.Errorf("no file path")
	}
	bytes, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, s)
}

func (s *Store) syncFromKVLocked() {
	kvBytes, err := LoadFromKV(kvStoreKey)
	if err == nil && len(kvBytes) > 0 {
		var temp struct {
			Projects       []Project       `json:"projects"`
			Experiences    []Experience    `json:"experiences"`
			Educations     []Education     `json:"educations"`
			Certifications []Certification `json:"certifications"`
		}
		if err := json.Unmarshal(kvBytes, &temp); err == nil && len(temp.Projects) > 0 {
			s.Projects = temp.Projects
			s.Experiences = temp.Experiences
			s.Educations = temp.Educations
			s.Certifications = temp.Certifications
			return
		}
	} else if err == nil && len(kvBytes) == 0 && len(s.Projects) > 0 {
		// First time: KV is empty, seed it with default data
		bytes, _ := json.MarshalIndent(s, "", "  ")
		if len(bytes) > 0 {
			_ = SaveToKV(kvStoreKey, bytes)
		}
	}
}

func (s *Store) saveToFile() error {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	// Persist to Cloudflare KV
	_ = SaveToKV(kvStoreKey, bytes)

	// Best-effort write to local file (in native Go / local environment)
	if s.filePath != "" {
		_ = os.WriteFile(s.filePath, bytes, 0644)
	}
	return nil
}

// Slugify generates a URL-safe ID from a string
func Slugify(text string) string {
	text = strings.ToLower(text)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug := reg.ReplaceAllString(text, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = "item"
	}
	return slug
}

// =========================================================================
// PROJECTS CRUD
// =========================================================================

func (s *Store) GetAllProjects(category string) []Project {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	if category == "" || strings.ToLower(category) == "all" {
		out := make([]Project, len(s.Projects))
		copy(out, s.Projects)
		return out
	}

	var filtered []Project
	catLower := strings.ToLower(category)
	for _, p := range s.Projects {
		if strings.Contains(strings.ToLower(p.Category), catLower) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (s *Store) GetProjectByID(id string) *Project {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for _, p := range s.Projects {
		if strings.EqualFold(p.ID, id) {
			copyProj := p
			return &copyProj
		}
	}
	return nil
}

func (s *Store) CreateProject(p Project) (*Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	// Auto-generate ID if empty
	if strings.TrimSpace(p.ID) == "" {
		baseSlug := Slugify(p.Title)
		p.ID = baseSlug
		// Ensure uniqueness
		count := 1
		for {
			collision := false
			for _, existing := range s.Projects {
				if existing.ID == p.ID {
					collision = true
					break
				}
			}
			if !collision {
				break
			}
			count++
			p.ID = fmt.Sprintf("%s-%d", baseSlug, count)
		}
	}

	// Auto-generate sequential number if empty e.g. #05
	if strings.TrimSpace(p.Number) == "" {
		p.Number = fmt.Sprintf("#%02d", len(s.Projects)+1)
	}

	s.Projects = append(s.Projects, p)
	_ = s.saveToFile()
	return &p, nil
}

func (s *Store) UpdateProject(id string, update Project) (*Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for idx, p := range s.Projects {
		if strings.EqualFold(p.ID, id) {
			// Retain original ID and number if not specified
			if strings.TrimSpace(update.ID) == "" {
				update.ID = p.ID
			}
			if strings.TrimSpace(update.Number) == "" {
				update.Number = p.Number
			}
			s.Projects[idx] = update
			_ = s.saveToFile()
			return &update, nil
		}
	}
	return nil, fmt.Errorf("project with ID '%s' not found", id)
}

func (s *Store) DeleteProject(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	foundIdx := -1
	for idx, p := range s.Projects {
		if strings.EqualFold(p.ID, id) {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return fmt.Errorf("project with ID '%s' not found", id)
	}

	s.Projects = append(s.Projects[:foundIdx], s.Projects[foundIdx+1:]...)
	_ = s.saveToFile()
	return nil
}

// NextProjectSequence returns the next available sequence number and suggested ID
func (s *Store) NextProjectSequence() (string, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()
	nextNum := len(s.Projects) + 1
	return fmt.Sprintf("#%02d", nextNum), nextNum
}

// =========================================================================
// CERTIFICATIONS CRUD
// =========================================================================

func (s *Store) GetAllCertifications(category string) []Certification {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	if category == "" || strings.ToLower(category) == "all" {
		out := make([]Certification, len(s.Certifications))
		copy(out, s.Certifications)
		return out
	}

	var filtered []Certification
	catLower := strings.ToLower(category)
	for _, c := range s.Certifications {
		if strings.EqualFold(c.Category, catLower) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (s *Store) GetCertificationByID(id int) *Certification {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for _, c := range s.Certifications {
		if c.ID == id {
			copyCert := c
			return &copyCert
		}
	}
	return nil
}

func (s *Store) CreateCertification(c Certification) (*Certification, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	// Auto-increment ID if 0 or empty
	if c.ID <= 0 {
		maxID := 0
		for _, existing := range s.Certifications {
			if existing.ID > maxID {
				maxID = existing.ID
			}
		}
		c.ID = maxID + 1
	}

	s.Certifications = append(s.Certifications, c)
	_ = s.saveToFile()
	return &c, nil
}

func (s *Store) UpdateCertification(id int, update Certification) (*Certification, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for idx, c := range s.Certifications {
		if c.ID == id {
			update.ID = id
			s.Certifications[idx] = update
			_ = s.saveToFile()
			return &update, nil
		}
	}
	return nil, fmt.Errorf("certification with ID %d not found", id)
}

func (s *Store) DeleteCertification(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	foundIdx := -1
	for idx, c := range s.Certifications {
		if c.ID == id {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return fmt.Errorf("certification with ID %d not found", id)
	}

	s.Certifications = append(s.Certifications[:foundIdx], s.Certifications[foundIdx+1:]...)
	_ = s.saveToFile()
	return nil
}

func (s *Store) NextCertificationID() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	maxID := 0
	for _, c := range s.Certifications {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	return maxID + 1
}

// =========================================================================
// EXPERIENCES & EDUCATIONS CRUD
// =========================================================================

func (s *Store) GetAllExperiences() []Experience {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()
	out := make([]Experience, len(s.Experiences))
	copy(out, s.Experiences)
	return out
}

func (s *Store) CreateExperience(e Experience) (*Experience, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	if strings.TrimSpace(e.ID) == "" {
		e.ID = fmt.Sprintf("exp-%d", len(s.Experiences)+1)
	}

	s.Experiences = append(s.Experiences, e)
	_ = s.saveToFile()
	return &e, nil
}

func (s *Store) UpdateExperience(id string, update Experience) (*Experience, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for idx, e := range s.Experiences {
		if strings.EqualFold(e.ID, id) {
			update.ID = id
			s.Experiences[idx] = update
			_ = s.saveToFile()
			return &update, nil
		}
	}
	return nil, fmt.Errorf("experience with ID '%s' not found", id)
}

func (s *Store) DeleteExperience(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	foundIdx := -1
	for idx, e := range s.Experiences {
		if strings.EqualFold(e.ID, id) {
			foundIdx = idx
			break
		}
	}
	if foundIdx == -1 {
		return fmt.Errorf("experience with ID '%s' not found", id)
	}

	s.Experiences = append(s.Experiences[:foundIdx], s.Experiences[foundIdx+1:]...)
	_ = s.saveToFile()
	return nil
}

func (s *Store) GetAllEducations() []Education {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()
	out := make([]Education, len(s.Educations))
	copy(out, s.Educations)
	return out
}

func (s *Store) CreateEducation(ed Education) (*Education, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	if strings.TrimSpace(ed.ID) == "" {
		ed.ID = fmt.Sprintf("edu-%d", len(s.Educations)+1)
	}

	s.Educations = append(s.Educations, ed)
	_ = s.saveToFile()
	return &ed, nil
}

func (s *Store) UpdateEducation(id string, update Education) (*Education, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	for idx, ed := range s.Educations {
		if strings.EqualFold(ed.ID, id) {
			update.ID = id
			s.Educations[idx] = update
			_ = s.saveToFile()
			return &update, nil
		}
	}
	return nil, fmt.Errorf("education with ID '%s' not found", id)
}

func (s *Store) DeleteEducation(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFromKVLocked()

	foundIdx := -1
	for idx, ed := range s.Educations {
		if strings.EqualFold(ed.ID, id) {
			foundIdx = idx
			break
		}
	}
	if foundIdx == -1 {
		return fmt.Errorf("education with ID '%s' not found", id)
	}

	s.Educations = append(s.Educations[:foundIdx], s.Educations[foundIdx+1:]...)
	_ = s.saveToFile()
	return nil
}
