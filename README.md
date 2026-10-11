# Helpdesk API - Go Fiber Backend

Backend REST API untuk sistem manajemen layanan bantuan (*Helpdesk*) yang dibangun menggunakan **Go**, **Fiber v2**, **PostgreSQL (pgx/v5)**, serta menerapkan **JWT Authentication** dan **Role-Based Access Control (RBAC)**.

## 🚀 Fitur Utama
- **Authentication**: Registrasi, Login (dengan validasi *bcrypt*), dan pengelolaan Token JWT (`golang-jwt/jwt/v5`).
- **Authorization & RBAC**: Pembatasan hak akses endpoint berdasarkan peran pengguna (*admin* dan *user*).
- **Ownership Validation**: Memastikan pengguna hanya dapat mengakses atau memodifikasi data miliknya sendiri (kecuali memiliki hak akses administratif).
- **Database Pooling**: Koneksi efisien menggunakan `pgxpool`.
- **Structured Logging**: Menggunakan `log/slog` bawaan Go untuk pencatatan request HTTP.

## 🛠️ Tech Stack
- **Language**: Go (Golang)
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **Driver/Pool**: `pgx/v5`
- **Security**: `golang-jwt/jwt/v5`, `bcrypt`

## ⚙️ Cara Instalasi & Menjalankan Proyek

1. **Clone repository ini:**
   ```bash
   git clone [https://github.com/mrcleverloki/helpdesk-api.git](https://github.com/mrcleverloki/helpdesk-api.git)
   cd helpdesk-api
Konfigurasi Database PostgreSQL:
Buat database baru dengan nama helpdesk_db:

SQL
CREATE DATABASE helpdesk_db;
Konfigurasi Environment Variable (.env):
Buat file .env di root direktori dan sesuaikan konfigurasinya:

Cuplikan kode
PORT=3000
DATABASE_URL=postgres://postgres:postgres@localhost:5432/helpdesk_db?sslmode=disable
JWT_SECRET=your_super_secret_key_here
Jalankan Aplikasi:

Go
go run .
🧪 Pengujian Endpoint
Pengujian dapat dilakukan menggunakan alat seperti Thunder Client atau Postman:

POST /api/v1/auth/login - Masuk ke sistem untuk mendapatkan access_token.

GET /api/v1/auth/me - Mendapatkan profil user dan daftar hak akses (permissions).

GET /api/v1/users - Menampilkan daftar user (Khusus Admin).

GET /api/v1/users/:id - Menampilkan detail user berdasarkan ID (Menerapkan aturan ownership).

👤 Author
Achmad Derajat Wibowo Tamyis (mrcleverloki)
