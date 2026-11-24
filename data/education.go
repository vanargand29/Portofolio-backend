package data

type Education struct {
	ID          string   `json:"id"`
	Date        string   `json:"date"`
	Title       string   `json:"title"`
	Institution string   `json:"institution"`
	Major       string   `json:"major"`
	Description string   `json:"description"`
	Details     string   `json:"details"`
	GPA         string   `json:"gpa"`
	Tags        []string `json:"tags"`
	CertHref    string   `json:"certHref"`
	IconType    string   `json:"iconType"`
}

var Educations = []Education{
	{
		ID:          "unej-s1-informatika",
		Date:        "2019 - 2026",
		Title:       "Universitas Jember",
		Institution: "Universitas Jember",
		Major:       "S1 Informatika",
		Description: "Fakultas Ilmu Komputer · Jurusan Informatika",
		Details:     "Menempuh pendidikan Sarjana Informatika dengan konsentrasi Software Engineering dan Web Technologies. Menguasai fondasi algoritma, struktur data, kecerdasan buatan, sistem pakar, serta metodologi pengembangan perangkat lunak modern. Berperan aktif dalam kegiatan kemahasiswaan dan riset sistem skrining kesehatan.",
		GPA:         "IPK: 3.66",
		Tags:        []string{"Software Engineering", "Algorithms", "Expert Systems", "Web Tech"},
		IconType:    "academic",
	},
}

func GetAllEducations() []Education {
	return Educations
}
