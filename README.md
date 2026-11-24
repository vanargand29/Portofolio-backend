# Portfolio Backend API (Golang & Cloudflare Workers)

Backend RESTful API untuk portofolio yang dibangun menggunakan **Golang** dan siap dideploy ke **Cloudflare Workers** menggunakan teknologi **WebAssembly (Wasm)**.

Backend ini juga mendukung mode **Local Development Standalone** (bisa dijalankan langsung dengan `go run .` tanpa perlu emulator Cloudflare).

---

## 📁 Struktur Direktori

```text
portofolio-backend/
├── data/
│   ├── projects.go          # Data lengkap 4 project portofolio
│   ├── experience.go        # Data pengalaman kerja & magang (Ruangguru, Dicoding)
│   ├── education.go         # Data pendidikan (Universitas Jember S1 Informatika)
│   └── certifications.go   # Data 33 sertifikat & kategori
├── handlers/
│   ├── health.go            # Health check & dokumentasi endpoint
│   ├── projects.go          # Handler endpoint Projects
│   └── resume.go            # Handler endpoint Experience, Education, Certifications, Resume
├── middleware/
│   └── cors.go              # CORS middleware (mendukung localhost & Netlify)
├── build/                   # Asset Wasm & runtime Cloudflare Workers (auto-generated)
├── go.mod                   # Definisi Go module & dependency
├── go.sum                   # Checksum dependency Go
├── main.go                  # Entry point aplikasi & routing
├── package.json             # Build script & dependency wrangler
├── wrangler.jsonc           # Konfigurasi Cloudflare Workers
└── README.md                # Dokumentasi & panduan penggunaan
```

---

## 🚀 Daftar Endpoint API

Semua response berformat JSON standar dengan metadata status `success` atau `error`.

| Method | Endpoint | Deskripsi | Query Parameters (Opsional) |
|---|---|---|---|
| `GET` | `/` | Health check, API sitemap, dan server info | - |
| `GET` | `/api/projects` | Mengambil seluruh daftar project | `?category=all` / `fullstack` / `ml` |
| `GET` | `/api/projects/{id}` | Mengambil detail 1 project spesifik | Contoh: `/api/projects/jstore` |
| `GET` | `/api/resume` | Mengambil seluruh data resume gabungan | - |
| `GET` | `/api/experience` | Mengambil daftar pengalaman kerja | - |
| `GET` | `/api/education` | Mengambil data riwayat pendidikan | - |
| `GET` | `/api/certifications`| Mengambil daftar 33 sertifikat | `?category=web` / `mobile` / `soft` / `seminar` / `organization` / `lomba` |

---

## 💻 1. Cara Menjalankan Secara Lokal (Local Run)

Anda memiliki dua pilihan untuk menjalankan di komputer lokal:

### Opsi A: Menggunakan Native Go (Paling Cepat & Ringan)
Buka terminal di folder `D:\JAYA\Code\Porto\portofolio-backend`:
```bash
go run .
```
API akan langsung berjalan di:
👉 **`http://localhost:8080`**

Coba buka di browser atau Postman:
- `http://localhost:8080/`
- `http://localhost:8080/api/projects`
- `http://localhost:8080/api/resume`

---

### Opsi B: Menggunakan Emulator Cloudflare Workers (Wrangler Dev)
Jika ingin menguji lingkungan simulasi Cloudflare Workers lokal:
```bash
# 1. Install wrangler jika belum
npm install

# 2. Jalankan emulator lokal
npm run dev
```
Wrangler akan otomatis melakukan kompilasi Go ke WebAssembly (`app.wasm`) dan menjalankan simulasi di `http://localhost:8787`.

---

## ☁️ 2. Cara Deploy ke Cloudflare Workers

Aplikasi ini sudah dikonfigurasi untuk Cloudflare Workers menggunakan WebAssembly (`workers-go`).

### Langkah 1: Login ke Akun Cloudflare
Di terminal Anda, jalankan:
```bash
npx wrangler login
```
Browser akan terbuka otomatis. Klik tombol **Authorize / Allow** untuk memberikan izin akses ke akun Cloudflare Anda.

### Langkah 2: Deploy ke Cloudflare
Cukup jalankan perintah:
```bash
npm run deploy
```
*Proses ini akan otomatis:*
1. Meng-compile Go menjadi WebAssembly (`build/app.wasm`).
2. Menghasilkan file wrapper V8 JavaScript (`build/worker.mjs`).
3. Mengunggah Wasm binary ke Cloudflare Workers edge network.

Setelah selesai, terminal akan menampilkan URL public API Anda, contoh:
```text
Uploaded portfolio-backend-api (1.23 sec)
Deployed portfolio-backend-api triggers
  https://portfolio-backend-api.<username>.workers.dev
```

---

## 🔗 3. Menghubungkan Backend ke Frontend React (Netlify)

Di project Frontend React Anda (`D:\JAYA\Code\Porto\Frontend React`):

1. Buat / edit file `.env` di root project React:
   ```env
   # Saat pengembangan lokal:
   VITE_API_BASE_URL=http://localhost:8080

   # Setelah backend dideploy ke Cloudflare:
   # VITE_API_BASE_URL=https://portfolio-backend-api.<username>.workers.dev
   ```

2. Di dashboard Netlify (agar saat di-deploy online otomatis mengambil data dari Cloudflare):
   - Buka: **Site configuration** -> **Environment variables**.
   - Tambahkan variabel:
     - **Key**: `VITE_API_BASE_URL`
     - **Value**: `https://portfolio-backend-api.<username>.workers.dev`

3. Contoh fetch di komponen React:
   ```javascript
   const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

   // Mengambil data project
   useEffect(() => {
     fetch(`${API_BASE}/api/projects`)
       .then(res => res.json())
       .then(response => {
         if (response.success) {
           setProjects(response.data);
         }
       })
       .catch(err => console.error("Gagal mengambil data:", err));
   }, []);
   ```

---

## 🛡️ CORS & Keamanan
CORS sudah otomatis dikonfigurasi di `middleware/cors.go`:
- Menerima request dari `http://localhost:5173`, `http://localhost:3000`, dan domain `*.netlify.app`.
- Menangani HTTP `OPTIONS` preflight request dengan status `204 No Content`.
