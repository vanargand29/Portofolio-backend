package data

import "strings"

type Certification struct {
	ID         int    `json:"id"`
	Category   string `json:"category"`
	Title      string `json:"title"`
	Issuer     string `json:"issuer"`
	Href       string `json:"href"`
	IsVerified bool   `json:"isVerified"`
}

var Certifications = []Certification{
	{
		ID:         1,
		Category:   "web",
		Title:      "Belajar Prinsip Pemrograman SOLID",
		Issuer:     "Dicoding · Sep 2021 - Sep 2024",
		Href:       "https://drive.google.com/file/d/1VRFvSJX5APotzU-_9D_mEFUt5u0bZ5wc/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         2,
		Category:   "web",
		Title:      "Memulai Dasar Pemrograman untuk Menjadi Pengembang Software",
		Issuer:     "Dicoding · Agu 2021 - Agu 2024",
		Href:       "https://drive.google.com/file/d/1JhGlB35VcKGWDbACk3QzuZfwuLwzQapT/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         3,
		Category:   "web",
		Title:      "Meniti Karier sebagai Software Developer",
		Issuer:     "Dicoding · Nov 2021 - Nov 2024",
		Href:       "https://drive.google.com/file/d/1HCuuSX4LltX2cWZfHTWx1t5tGyaWIW4p/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         4,
		Category:   "web",
		Title:      "Pengenalan ke Logika Pemrograman (Programming Logic 101)",
		Issuer:     "Dicoding · Agu 2021 - Agu 2024",
		Href:       "https://drive.google.com/file/d/1LixSChwz9x03v6I027aSWIbXeEI06-gL/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         5,
		Category:   "web",
		Title:      "Belajar Dasar Git dengan GitHub",
		Issuer:     "Dicoding · Agu 2021 - Agu 2024",
		Href:       "https://drive.google.com/file/d/1Oe6B5Be3PeGaJZL19nAZKFK6cV5FAclI/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         6,
		Category:   "web",
		Title:      "Memulai Pemrograman dengan Java",
		Issuer:     "Dicoding · Sep 2021 - Sep 2024",
		Href:       "https://drive.google.com/file/d/1v0neI6kmb9OrwpYVjXPTKkC6OiyYq-pU/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         7,
		Category:   "web",
		Title:      "Memulai Pemrograman dengan C",
		Issuer:     "Dicoding · Nov 2021 - Nov 2024",
		Href:       "https://drive.google.com/file/d/1BVYlhSuJLyXdw9QmX8X_XL0uA0xbSq-J/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         8,
		Category:   "web",
		Title:      "Data Wrangling Python",
		Issuer:     "DQLab · Sep 2020",
		Href:       "https://drive.google.com/file/d/1qQ4nTNk8F7ycVaFmIb1CBxd4CIEigFk1/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         9,
		Category:   "web",
		Title:      "Python Fundamental for Data Science",
		Issuer:     "DQLab · Sep 2020",
		Href:       "https://drive.google.com/file/d/1Bj0zoJrEh3y4msDdMQuuQdKyoDtbfJv-/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         10,
		Category:   "web",
		Title:      "Introduction to Data Science with Python",
		Issuer:     "DQLab · Sep 2020",
		Href:       "https://drive.google.com/file/d/1mV0MKtpWvwzO0M2TIG2xUfozfHG_jJ__/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         11,
		Category:   "mobile",
		Title:      "Belajar Fundamental Aplikasi Android",
		Issuer:     "Dicoding · Nov 2021 - Nov 2024",
		Href:       "https://drive.google.com/file/d/1DYjhbRzIbPWvSJXVwGB58BudZIDHhtKA/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         12,
		Category:   "mobile",
		Title:      "Belajar Membuat Aplikasi Android untuk Pemula",
		Issuer:     "Dicoding · Okt 2021 - Okt 2024",
		Href:       "https://drive.google.com/file/d/10jI_CR6JIFi9EeGPCbQS0PDhqvWKKynp/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         13,
		Category:   "mobile",
		Title:      "Memulai Pemrograman dengan Kotlin",
		Issuer:     "Dicoding · Sep 2021 - Sep 2024",
		Href:       "https://drive.google.com/file/d/1nq8K1_KJ1TL-4KwzBdhvy4itGSNR3Lq2/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         14,
		Category:   "mobile",
		Title:      "Belajar Membuat Aplikasi Flutter untuk Pemula",
		Issuer:     "Dicoding · Des 2021 - Des 2024",
		Href:       "https://drive.google.com/file/d/1X7FG1pKasimesa9TXRkw65YkyDqcOntu/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         15,
		Category:   "mobile",
		Title:      "Memulai Pemrograman dengan Dart",
		Issuer:     "Dicoding · Okt 2021 - Okt 2024",
		Href:       "https://drive.google.com/file/d/18HU5A7xkrqFwWSVeOqjkQx8n_oCSM1zI/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         16,
		Category:   "mobile",
		Title:      "Belajar Fundamental Aplikasi Flutter",
		Issuer:     "Dicoding · Jan 2022 - Jan 2025",
		Href:       "https://drive.google.com/file/d/1fSXTTTgbg941SOylZNneZLFyDF3Ow66x/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         17,
		Category:   "soft",
		Title:      "Certificate of Casual Conversation",
		Issuer:     "Dicoding",
		Href:       "https://drive.google.com/file/d/1MYHPy-gIxt-CEFRPrYRcMSapEDrw_W7J/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         18,
		Category:   "soft",
		Title:      "Peningkatan Produktivitas dan Peluang Berwirausaha di Masa Covid-19",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/175LJwm6Xr-BGTvJgb4zDWCN07qiUwlVU/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         19,
		Category:   "soft",
		Title:      "Seminar Bisnis Konsultasi Mahasiswa Preneur",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/1FIP4NG2p1jUzS-RrwGXUOMebF14knFOF/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         20,
		Category:   "soft",
		Title:      "Mengenali Diri Bekal Mendasar dalam Beradaptasi",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/1jg2CC2ARRswOiKuFDjmVESx_-bp3JHBn/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         21,
		Category:   "seminar",
		Title:      "AWS x Dicoding LIVE: Storage dan Database di AWS",
		Issuer:     "Dicoding",
		Href:       "https://drive.google.com/file/d/1k1kHb9XgD_PlxNRfE-NhJLgmsLPZYmEI/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         22,
		Category:   "seminar",
		Title:      "SPOGISFY (Sharing of GIS Laboratory)",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/1mrAlZqu4eGOSRyt3s1UibYBgJ22axAN8/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         23,
		Category:   "seminar",
		Title:      "YOTClass Design (Create Your Idea With CorelDraw)",
		Issuer:     "Workshop",
		Href:       "https://drive.google.com/file/d/1nH_M1WofuvZs_hgqAxocsVi7_fHgfIwq/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         24,
		Category:   "seminar",
		Title:      "POSTGRE (Pemro Share Talk & Greet)",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/1Y4eYLpQLebnkgKGxJeuoUnyim6Gr-C0j/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         25,
		Category:   "seminar",
		Title:      "ICOMITEE 2021 dan Workshop Bedah Artikel Ilmiah",
		Issuer:     "Workshop",
		Href:       "https://drive.google.com/file/d/15-8LCCVdOQY5VR8JeE5R9NF94TuBckLX/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         26,
		Category:   "seminar",
		Title:      "Compiler Algorithm Training",
		Issuer:     "Training",
		Href:       "https://drive.google.com/file/d/1S6IPfc5RJjWJaOEo_kNPY1R3UXXkyRyj/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         27,
		Category:   "seminar",
		Title:      "Upaya Restorasi oleh Pemerintah dalam Menghadapi Dampak Covid-19",
		Issuer:     "Seminar",
		Href:       "https://drive.google.com/file/d/1r1M3nB1EY2CiGeFYXW0qCItq6j64StTI/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         28,
		Category:   "seminar",
		Title:      "Weekend Tech Talk (Let's Learn Brainstorming Technique)",
		Issuer:     "Tech Talk",
		Href:       "https://drive.google.com/file/d/16NQhmK1qGERSjmfrVpgO9nG7AWuSZ5_6/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         29,
		Category:   "organization",
		Title:      "Sie Perlengkapan - Pembinaan dan Pengembangan Mahasiswa Baru (PPMB)",
		Issuer:     "Universitas Jember",
		Href:       "https://drive.google.com/file/d/1vFxXB_0YLc1sN1OAbykdVWlZq7zO3d6V/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         30,
		Category:   "organization",
		Title:      "Sie Perlengkapan - Friendship Competition 2021",
		Issuer:     "Universitas Jember",
		Href:       "https://drive.google.com/file/d/1NekwzDWqm6_Ju-q5UJzpI2-U2mj8OeAe/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         31,
		Category:   "organization",
		Title:      "Sie Acara - Pengenalan Kehidupan Kampus bagi Mahasiswa Baru (PKKMB)",
		Issuer:     "Universitas Jember",
		Href:       "https://drive.google.com/file/d/1Nteu6F9xOwutVoRPCjWbevjxg1PxSdhV/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         32,
		Category:   "lomba",
		Title:      "Capture The Flag (CTF)",
		Issuer:     "Competition",
		Href:       "https://drive.google.com/file/d/1s_Gb1Zb6FyhTLMeMOukq9htpXJ_hSvQX/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         33,
		Category:   "lomba",
		Title:      "Laos Arena",
		Issuer:     "Competition",
		Href:       "https://drive.google.com/file/d/14fSKWLw2nZOXoByNRiOhvF-L5r9Ll-EU/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         34,
		Category:   "lomba",
		Title:      "Competitive Programming Competition",
		Issuer:     "Competition",
		Href:       "https://drive.google.com/file/d/1bszmDafdBi2UCQ-FXZPWL77l7q6B5BMQ/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         35,
		Category:   "lomba",
		Title:      "Festival Menulis Surat Nasional",
		Issuer:     "Competition",
		Href:       "https://drive.google.com/file/d/1FiCoTd1jcPXWHtXCxQ-dw-G_sHBrYAvf/view?usp=sharing",
		IsVerified: true,
	},
	{
		ID:         36,
		Category:   "lomba",
		Title:      "Menulis puisi bertema Cinta",
		Issuer:     "Competition",
		Href:       "https://drive.google.com/file/d/1WL_ldvceOhM6bzpgVwN3lWlk-3l4D21l/view?usp=sharing",
		IsVerified: true,
	},
}

func GetAllCertifications(category string) []Certification {
	if category == "" || strings.ToLower(category) == "all" {
		return Certifications
	}
	filtered := make([]Certification, 0)
	target := strings.ToLower(category)
	for _, c := range Certifications {
		if strings.ToLower(c.Category) == target {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
