package data

import "strings"

type Project struct {
	ID             string   `json:"id"`
	Number         string   `json:"number"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	Category       string   `json:"category"`
	HeaderCategory string   `json:"headerCategory"`
	HeaderType     string   `json:"headerType"`
	UrlMock        string   `json:"urlMock"`
	Description    string   `json:"description"`
	Features       []string `json:"features"`
	Tech           []string `json:"tech"`
	Image          string   `json:"image"`
	Github         string   `json:"github"`
	IsPrivate      bool     `json:"isPrivate"`
}

var Projects = []Project{
	{
		ID:             "jstore",
		Number:         "#01",
		Title:          "J-Store – Game Currency Store Platform",
		Subtitle:       "High Performance Gaming Store Frontend & Checkout Flow",
		Category:       "Website",
		HeaderCategory: "GAME STORE",
		HeaderType:     "WEB APPLICATION",
		UrlMock:        "jstore-gaming.app",
		Description:    "J-Store adalah platform e-commerce mata uang game modern untuk game populer seperti Genshin Impact. Platform ini menyediakan katalog produk yang responsif, visual interaktif, kalkulasi harga otomatis, dan alur checkout yang aman serta siap integrasi payment gateway.",
		Features: []string{
			"Katalog produk mata uang game dengan filtering kategori instan",
			"Kalkulasi konversi mata uang dan diskon secara real-time",
			"Desain antarmuka mobile-first responsif di semua resolusi layar",
			"Struktur kode modular berbasis standar HTML5, CSS3, dan ES6+",
		},
		Tech:      []string{"HTML", "CSS", "JavaScript"},
		Image:     "/img/jstore(tampilan).webp",
		Github:    "#",
		IsPrivate: true,
	},
	{
		ID:             "tuberculosis",
		Number:         "#02",
		Title:          "Tuberculosis Expert System – AI Disease Screening",
		Subtitle:       "Clinical Diagnostic Inference Engine via Dempster-Shafer",
		Category:       "Machine Learning",
		HeaderCategory: "HEALTHCARE",
		HeaderType:     "EXPERT SYSTEM",
		UrlMock:        "tb-expertsystem.health.id",
		Description:    "Aplikasi sistem pakar skrining gejala Tuberkulosis berbasis web menggunakan inferensi Dempster-Shafer. Menganalisis multi-gejala klinis dan menghitung probabilitas diagnosis secara otomatis untuk membantu keputusan tenaga kesehatan secara akurat.",
		Features: []string{
			"Form input gejala klinis interaktif dengan pembobotan evidensial",
			"Algoritma inferensi kombinasi Dempster-Shafer secara otomatis",
			"Visualisasi persentase probabilitas dan confidence score diagnosis (94.6%)",
			"Rekomendasi penanganan medis dan saran pemeriksaan sputum AFB lanjutan",
		},
		Tech:      []string{"Laravel", "MySQL", "Dempster-Shafer"},
		Image:     "/img/tuberculosis-preview.svg",
		Github:    "#",
		IsPrivate: true,
	},
	{
		ID:             "portfolio",
		Number:         "#03",
		Title:          "Personal Portfolio – Interactive Showcase",
		Subtitle:       "Modern React SPA with Smooth Horizontal Sliding Navigation",
		Category:       "Website",
		HeaderCategory: "PORTFOLIO",
		HeaderType:     "PERSONAL WEBSITE",
		UrlMock:        "arisandi-portfolio.netlify.app",
		Description:    "Website portfolio personal berkinerja tinggi dengan navigasi horizontal slider yang mulus, tema dark glassmorphism modern, micro-interactions, modal interaktif, dan filter sertifikasi dinamis.",
		Features: []string{
			"Horizontal snapping track dengan transisi halus dan wheel easing",
			"Desain modern glassmorphism dengan aksen pendaran merah khas",
			"Modal showcase interaktif dengan zoom screenshot beresolusi tinggi",
			"Timeline pengalaman, pendidikan, dan sertifikasi dengan multi-filter",
			"Formulir kontak terintegrasi mailto siap kirim",
		},
		Tech:      []string{"React", "Responsive UI", "HTML/CSS/JS"},
		Image:     "/img/porto(tampilan).webp",
		Github:    "https://github.com/vanargand29/Portofolio",
		IsPrivate: false,
	},
	{
		ID:             "backendApi",
		Number:         "#04",
		Title:          "Backend API & Admin CRUD – Content Manager",
		Subtitle:       "Headless REST API with Modern Administrative Management",
		Category:       "Backend",
		HeaderCategory: "BACKEND API",
		HeaderType:     "REST API & ADMIN PANEL",
		UrlMock:        "api.arisandi.workers.dev",
		Description:    "Arsitektur REST API lengkap dengan dashboard admin modern untuk manajemen data portfolio secara dinamis. Mendukung operasi CRUD penuh pada data proyek dan sertifikasi dengan endpoint terstruktur dan aman.",
		Features: []string{
			"REST API endpoint terstruktur untuk manajemen proyek dan resume",
			"Panel admin intuitif untuk operasi CRUD data secara real-time",
			"Format pertukaran data JSON modular dan terorganisir",
			"Deployment cepat dan serverless via Cloudflare Workers",
			"CORS middleware terkonfigurasi untuk komunikasi domain aman",
		},
		Tech:      []string{"Go", "Node.js", "Express", "Cloudflare Workers"},
		Image:     "/img/adminAPI.webp",
		Github:    "#",
		IsPrivate: true,
	},
}

func GetAllProjects(category string) []Project {
	if category == "" || strings.ToLower(category) == "all" {
		return Projects
	}
	filtered := make([]Project, 0)
	target := strings.ToLower(category)
	for _, p := range Projects {
		if strings.ToLower(p.Category) == target {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func GetProjectByID(id string) *Project {
	for _, p := range Projects {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
