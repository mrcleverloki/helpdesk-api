# Helpdesk API - Clean Architecture & Go Fiber

Backend REST API untuk sistem helpdesk yang dikembangkan sebagai bagian dari praktikum Pemrograman Backend Lanjut. Dibangun menggunakan bahasa **Go**, framework **Fiber**, dan database **PostgreSQL** dengan menerapkan pola **Clean Architecture**.

## 🛠 Tech Stack
- **Bahasa Pemrograman**: Go
- **Web Framework**: Fiber v2 (`github.com/gofiber/fiber/v2`)
- **Database & Driver**: PostgreSQL dengan `pgx/v5` (*connection pooling* via `pgxpool`)
- **Autentikasi & Keamanan**: 
  - JSON Web Token (`github.com/golang-jwt/jwt/v5`)
  - Password Hashing (`golang.org/x/crypto/bcrypt`)
  - Kriptografi & Token Acak (`crypto/rand`, SHA-256)
- **Logging & Konfigurasi**: `slog`, `lumberjack`, `godotenv`

## 📁 Struktur Proyek
```text
helpdesk-api/
├── app/
│   ├── model/       # Struktur data, request/response, dan entitas model
│   ├── repository/  # Lapisan akses database PostgreSQL (User & Refresh Token)
│   └── service/     # Lapisan business logic, validasi aturan, dan alur layanan
├── config/          # Pengaturan aplikasi, inisialisasi Fiber, env, dan error handler
├── database/        # Manajemen koneksi pool database PostgreSQL
├── helper/          # Fungsi bantu global (HTTP response, JWT, kriptografi, context)
├── middleware/      # Middleware keamanan (Autentikasi JWT, Rate Limiter, CORS, Header JSON)
├── migrations/      # Berkas migrasi skema SQL database
├── route/           # Definisi pengelompokan endpoint API
└── main.go          # Titik masuk utama aplikasi (App Wiring & Graceful Shutdown)
⚙️ Konfigurasi Environment (.env)
Buat file .env di root direktori berdasarkan contoh berikut:

Cuplikan kode
APP_NAME="Helpdesk API"
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=helpdesk_db
JWT_SECRET=your-secure-random-32-byte-hex-string-from-openssl
JWT_ISSUER=praktikum-backend
JWT_ACCESS_TTL_MINUTES=15
JWT_REFRESH_TTL_DAYS=7
ALLOWED_ORIGINS=http://localhost:5173
🚀 Cara Menjalankan Proyek
Clone repositori ini dan pastikan Go sudah terinstal.

Jalankan unduhan dependensi:

Bash
go mod tidy
Jalankan migrasi database PostgreSQL:

Bash
psql -U postgres -d helpdesk_db -f migrations/001_init.sql
psql -U postgres -d helpdesk_db -f migrations/002_auth.sql
Jalankan aplikasi:

Bash
go run .
📌 Progres Modul
Modul 1-3: Inisialisasi proyek, arsitektur dasar, dan koneksi database PostgreSQL.

Modul 4: Implementasi Clean Architecture penuh, Structured Logging (slog), dan App Wiring.

Modul 5: Autentikasi & Keamanan meliputi Registrasi, Login, manajemen JWT Access Token, Refresh Token dengan rotasi token aman, middleware proteksi, dan Rate Limiter.
