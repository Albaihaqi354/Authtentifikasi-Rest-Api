# Authtentifikasi Rest API

Authtentifikasi Rest API adalah proyek yang menyediakan endpoint RESTful untuk autentikasi pengguna. Proyek ini menggunakan teknologi seperti Node.js, Express, dan JSON Web Tokens (JWT) untuk memastikan keamanan dan kecepatan dalam proses autentikasi.

## Fitur

- **Registrasi Pengguna:** Pengguna dapat membuat akun baru dengan username dan password.
- **Login Pengguna:** Pengguna dapat masuk menggunakan username dan password mereka.
- **JWT Authentication:** Setelah login berhasil, pengguna akan menerima token JWT yang harus digunakan untuk mengakses endpoint yang memerlukan autentikasi.
- **Proteksi Rute:** Rute tertentu hanya dapat diakses jika pengguna memiliki token JWT yang valid.

## Prasyarat

Pastikan Anda telah menginstal versi terbaru dari Node.js dan npm di lingkungan Anda.

## Instalasi

1. Clone repository ini ke mesin lokal Anda:

   ```bash
   git clone https://github.com/Albaihaqi354/Authtentifikasi-Rest-Api.git

2. Masuk ke direktori proyek:

   ```bash
   cd Authtentifikasi-Rest-Api

3. Instal dependensi yang diperlukan:

   ```bash
   npm install

## Konfigurasi

Buat file .env di root proyek dan tambahkan variabel-variabel berikut:

   ```bash
    PORT=5000
    MONGO_URI=mongodb://localhost:27017/authtentifikasi_db
    JWT_SECRET=your_jwt_secret
