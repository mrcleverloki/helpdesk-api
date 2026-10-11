# AI Usage - Helpdesk API

Dokumen ini merangkum bentuk kolaborasi dan bantuan AI (Gemini) selama pengembangan proyek *Helpdesk API* secara keseluruhan.

## Cakupan Bantuan AI dalam Proyek
1. **Inisialisasi & Arsitektur Proyek**
   - Membantu merancang struktur folder berbasis Clean/Layered Architecture (`repository`, `service`, `handler`, `route`) menggunakan framework Go Fiber.
   - Menyusun konfigurasi koneksi database PostgreSQL menggunakan `pgxpool` dan penanganan *connection pooling*.

2. **Autentikasi & Keamanan (Auth & RBAC)**
   - Membantu implementasi autentikasi berbasis JSON Web Tokens (JWT), *password hashing* dengan `bcrypt`, serta pembuatan middleware *Role-Based Access Control* (RBAC) dan pemeriksaan kepemilikan data (*ownership authorization*).

3. **Validasi & Penanganan Respons**
   - Mengintegrasikan validasi deklaratif menggunakan `validator/v10` beserta penyesuaian custom error messages agar lebih ramah pengguna.
   - Mendesain struktur respons standar (`WebResponse` dan `ErrorResponse`) dengan field fleksibel (`any`) guna mendukung berbagai bentuk data dan metadata.

4. **Pagination Tingkat Lanjut (Keyset Cursor Pagination)**
   - Membantu merancang logika *keyset cursor pagination* berbasis base64 encoding pada field `CreatedAt` dan `ID` untuk memastikan stabilitas performa data list yang besar tanpa duplikasi atau data yang terlewat.
   - Menyesuaikan query repository (`FindAllCursor`) dengan PostgreSQL.

## Batasan & Tanggung Jawab Pengembang
- Seluruh kode yang dihasilkan didiskusikan, di-debug, dan diuji secara mandiri oleh pengembang menggunakan *Thunder Client*.
- Seluruh logika bisnis, penyesuaian tipe data, dan integrasi akhir sepenuhnya divalidasi dan dikontrol secara manual oleh pengembang.