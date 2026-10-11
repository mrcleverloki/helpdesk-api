# Helpdesk API

Project API backend menggunakan bahasa Go dan framework Fiber v2 untuk memenuhi modul praktikum Pemrograman Backend Lanjut.

## Tech Stack
- Bahasa: Go (Golang)
- Framework: Fiber v2
- Database: In-memory (sementara, sebelum PostgreSQL di Modul 3)
- Tools: VS Code, Thunder Client, Git/GitHub

## API Endpoints (Modul 2)
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| GET | `/api/v1/health` | Cek status server |
| GET | `/api/v1/users` | Ambil daftar user (mendukung paginasi, search, sort, filter) |
| GET | `/api/v1/users/:id` | Ambil detail satu user berdasarkan ID |
| POST | `/api/v1/users` | Tambah user baru (status 201 Created + Header Location) |
| PUT | `/api/v1/users/:id` | Ganti seluruh data user (Replace) |
| PATCH | `/api/v1/users/:id` | Perbarui sebagian data user (Partial Update) |
| DELETE | `/api/v1/users/:id` | Hapus user (status 204 No Content) |
