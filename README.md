# Helpdesk API - Clean Architecture & Go Fiber

API backend layanan helpdesk yang dibangun menggunakan bahasa Go, framework Fiber, dan basis data PostgreSQL dengan menerapkan Clean Architecture.

## Fitur Utama
- Clean Architecture: Pemisahan layer antara Model, Repository, Service, Helper, Middleware, dan Route.
- Pure Business Rules: Logika validasi bisnis terisolasi dari HTTP context (fiber.Ctx).
- Structured Logging: Pencatatan log JSON dengan slog dan lumberjack.
- Database Connection Pooling: Menggunakan pgx/v5 untuk PostgreSQL.

## Struktur Direktori
- app/model: Entitas data, struct request, dan response.
- app/repository: Akses data dan query SQL.
- app/service: Logika bisnis dan HTTP handler.
- config: Konfigurasi env, Fiber app, dan logger.
- database: Inisialisasi pool PostgreSQL.
- helper: Fungsi bantu presenter dan request parser.
- middleware: Keamanan, CORS, dan validasi JSON.
- route: Pendaftaran endpoint API.

## Cara Menjalankan
1. Clone repository dan buat file .env sesuai konfigurasi lokal.
2. Jalankan migrasi database PostgreSQL.
3. Jalankan aplikasi dengan perintah:
   go run .
