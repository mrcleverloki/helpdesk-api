# Transparansi Penggunaan AI (AI-Usage)

Dokumen ini menjelaskan bagaimana asisten AI (Gemini) digunakan dalam proses pengembangan proyek **Helpdesk API**.

## Cakupan Bantuan AI
1. **Arsitektur & Boilerplate Code**: Membantu merancang struktur folder proyek berbasis *Clean Architecture* sederhana untuk Go Fiber serta perumusan struktur tabel PostgreSQL.
2. **Implementasi Keamanan**: Membantu menyusun logika hashing password menggunakan `bcrypt` serta pembuatan token JWT dengan *custom claims* (`Username`, `Role`).
3. **Debugging & Perbaikan Error**: Membantu menganalisis pesan error saat pengujian koneksi database, penyesuaian query SQL, serta validasi *password hash*.
4. **Dokumentasi**: Membantu menyusun file `README.md` dan dokumentasi penggunaan proyek secara terstruktur.

## Kontribusi Pengembang
Seluruh kode yang dihasilkan melalui konsultasi dengan AI telah melalui proses peninjauan (*review*), pengujian manual menggunakan Thunder Client, serta penyesuaian logika bisnis secara mandiri oleh pengembang (**Achmad Derajat Wibowo Tamyis**).
