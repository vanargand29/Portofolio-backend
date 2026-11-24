package data

type Experience struct {
	ID           string   `json:"id"`
	Date         string   `json:"date"`
	Title        string   `json:"title"`
	Organization string   `json:"organization"`
	Description  string   `json:"description"`
	Details      string   `json:"details"`
	Tags         []string `json:"tags"`
	CertHref     string   `json:"certHref"`
	IconType     string   `json:"iconType"`
}

var Experiences = []Experience{
	{
		ID:           "ruangguru-bootcamp",
		Date:         "Feb 2022 - Jul 2022",
		Title:        "Front-End Engineering Bootcamp",
		Organization: "Ruangguru",
		Description:  "Program Magang & Studi Independen Bersertifikat (MSIB) Front-End Engineering di PT Ruang Raya Indonesia.",
		Details:      "Kerja kolaboratif dengan Git & GitHub Version Control, Pemrograman modern Front-End JavaScript dan CSS, Component Programming dengan React, SSR Framework dengan Next.js, API Integration & Authentication, Identity Management, Performance Optimization, dan Dasar-dasar Project Management tim Scrum.",
		Tags:         []string{"React", "Next.js", "TypeScript", "REST API", "Cloud Deployment"},
		CertHref:     "https://drive.google.com/file/d/1_VLbu0URknhletqkEBRfFKFjgb0I5SWH/view?usp=sharing",
		IconType:     "code",
	},
	{
		ID:           "dicoding-multiplatform",
		Date:         "Agu 2021 - Jan 2022",
		Title:        "Pengembang Android & Multi-Platform",
		Organization: "Dicoding",
		Description:  "Program Magang & Studi Independen Bersertifikat (MSIB) Pengembang Aplikasi Android & Multi-Platform di PT Presentologics.",
		Details:      "Menyelesaikan kurikulum terakreditasi: Programming Logic 101, Git with GitHub, Prinsip Pemrograman SOLID, Pemrograman Android Native dengan Kotlin, Android Architecture Components, Pemrograman Multi-Platform dengan Dart & Flutter, serta implementasi fundamental state management.",
		Tags:         []string{"Android", "Kotlin", "Flutter", "Dart", "SOLID Principles"},
		CertHref:     "https://drive.google.com/file/d/1_VLbu0URknhletqkEBRfFKFjgb0I5SWH/view?usp=sharing",
		IconType:     "mobile",
	},
}

func GetAllExperiences() []Experience {
	return Experiences
}
