# Helpdesk API

Helpdesk API adalah backend layanan tiket bantuan (*helpdesk*) yang dibangun menggunakan bahasa pemrograman Go dengan framework **Fiber**, database **PostgreSQL**, dan arsitektur berlapis (Clean/Layered Architecture).

## Fitur Utama

1. **Clean & Modular Architecture**
   - Pemisahan tanggung jawab yang jelas antara `repository`, `service`, `helper`, `route`, dan `middleware`.
2. **Autentikasi & Otorisasi (RBAC)**
   - Autentikasi stateless menggunakan **JSON Web Tokens (JWT)**.
   - Keamanan sandi dengan **bcrypt** hashing.
   - *Role-Based Access Control* (RBAC) dan pemeriksaan kepemilikan data (*ownership authorization*).
3. **Validasi Deklaratif**
   - Menggunakan pustaka `validator/v10` untuk validasi input request yang bersih dan terstruktur disertai pesan error yang informatif.
4. **Keyset Cursor Pagination**
   - Penerapan paginasi berbasis cursor (menggunakan base64 encoding pada `CreatedAt` dan `ID`) untuk menjamin konsistensi dan performa tinggi pada list data yang besar tanpa masalah *offset skip*.
5. **Flexible Response Handling**
   - Standarisasi format respons JSON (`WebResponse`) yang mendukung data tunggal, daftar data, error, hingga metadata dinamis (*cursor/pagination meta*).

## Struktur Proyek

```text
helpdesk-api/
├── app/
│   ├── model/       # Definisi struct request, response, dan model data
│   ├── repository/  # Lapisan akses database PostgreSQL (pgxpool)
│   └── service/     # Lapisan logika bisnis
├── helper/          # Helper utilitas (JWT, validation, cursor, response)
├── middleware/      # Middleware Fiber (Auth, RBAC, JSON, Rate Limiter)
├── route/           # Pendaftaran rute API endpoint
├── database/        # Skema dan migrasi database
├── .env.example     # Contoh variabel lingkungan
├── AI-USAGE.md      # Catatan kolaborasi pengembangan berbasis AI
└── main.go          # Titik masuk utama aplikasi (Entrypoint)
Menjalankan Proyek
Salin file konfigurasi lingkungan:

Bash
cp .env.example .env
Sesuaikan konfigurasi database PostgreSQL pada file .env.

Jalankan aplikasi:

Bash
go run .
Endpoint Utama
Auth: /api/v1/auth/register, /api/v1/auth/login, /api/v1/auth/me

Users: /api/v1/users (Mendukung keyset cursor pagination via query parameter limit dan cursor)
