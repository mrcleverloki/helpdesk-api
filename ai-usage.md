# Transparansi Penggunaan AI (AI-USAGE)

Dokumen ini mencatat pembagian kerja antara pengembang (**Achmad Derajat Wibowo Tamyis**) dan asisten AI dalam pengembangan proyek **Helpdesk API**.

## Pembagian Tugas & Tanggung Jawab

### 1. Dikerjakan oleh Pengembang (User)
- Menentukan arsitektur dasar (*Clean Architecture*), struktur folder, serta perancangan skema database PostgreSQL (`users` dan `refresh_tokens`).
- Menulis, menguji, dan mengintegrasikan kode sumber secara bertahap menggunakan strategi *micro-commit*.
- Melakukan verifikasi keamanan, pengecekan error kompilasi (`go test`), serta memastikan seluruh alur autentikasi (register, login, JWT, rotasi refresh token) berjalan sesuai spesifikasi praktikum.

### 2. Peran AI sebagai Tools Bantu
- **Penyusunan Template & Kerangka Awal**: Membantu memberikan referensi struktur implementasi kode untuk *helper* kriptografi (`bcrypt`, SHA-256, JWT manager) dan *middleware*.
- **Debugging & Analisis Error**: Membantu mengidentifikasi akar permasalahan saat terjadi konflik tipe data, error package, atau ketidaksesuaian penandatanganan fungsi (*function signature*) pada lapisan *service* dan *repository*.
- **Penulisan Dokumentasi**: Membantu merumuskan penjelasan teknis dan format dokumentasi proyek.

Semua kode yang dihasilkan melalui bantuan AI telah dikurasi, dipahami secara mendalam, diuji coba langsung pada lingkungan lokal, dan disesuaikan sepenuhnya oleh pengembang.
