# Proyek Virtual Intern - Backend Service

Ini adalah proyek _backend service_ yang dibangun sebagai bagian dari program magang virtual. Proyek ini berfungsi sebagai REST API dasar yang dibuat menggunakan Go (Golang) dengan framework Fiber dan GORM sebagai ORM untuk berinteraksi dengan database MySQL.

## ✨ Fitur

- **REST API**: Menyediakan endpoint API yang terstruktur.
- **Framework Cepat**: Dibangun di atas [Fiber](https://gofiber.io/), sebuah framework web yang terinspirasi dari Express.js.
- **ORM (Object-Relational Mapping)**: Menggunakan [GORM](https://gorm.io/) untuk interaksi yang lebih mudah dan aman dengan database.
- **Struktur Proyek Bersih**: Mengadopsi struktur proyek yang terorganisir (mirip dengan _Clean Architecture_) untuk memisahkan _concern_ antara lapisan HTTP, _service_, dan _repository_.
- **Manajemen Konfigurasi**: Mengelola konfigurasi aplikasi (seperti kredensial database) menggunakan _environment variables_ (`.env`).

## 🛠️ Tumpukan Teknologi

- **Bahasa Pemrograman**: [Go (Golang)](https://golang.org/)
- **Framework Web**: [Fiber](https://gofiber.io/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: [MySQL](https://www.mysql.com/)
- **API Client**: [Postman](https://www.postman.com/)

## 📂 Struktur Proyek

.
├── cmd/
│ └── main.go # Entry point aplikasi
├── config/
│ └── config.go # Memuat konfigurasi dari .env
├── internal/ # Logika bisnis inti aplikasi
│ ├── builder/ # Builder untuk dependensi
│ ├── dto/ # Data Transfer Objects (untuk request/response)
│ ├── entity/ # Entitas/model domain
│ ├── http/ # Handler/controller HTTP
│ ├── repository/ # Lapisan akses data (berinteraksi dengan database)
│ └── service/ # Lapisan logika bisnis
├── pkg/
│ ├── database/ # Inisialisasi koneksi database
│ └── server/ # Inisialisasi server HTTP (Fiber)
├── .env.example # Contoh file environment
├── .gitignore
├── go.mod # File dependensi Go Modules
├── go.sum
└── evermos virtual intern.postman_collection.json # Koleksi Postman

## 🚀 Memulai Proyek

Berikut adalah langkah-langkah untuk menjalankan proyek ini di lingkungan lokal Anda.

### Prasyarat

- [Go](https://golang.org/dl/) versi 1.18 atau lebih tinggi
- [MySQL](https://dev.mysql.com/downloads/installer/)
- [Git](https://git-scm.com/downloads/)

### Instalasi & Konfigurasi

1.  **Clone repositori ini:**

    ```sh
    git clone <URL_REPOSITORI_ANDA>
    cd <NAMA_DIREKTORI_PROYEK>
    ```

2.  **Buat file `.env`:**
    Salin dari contoh yang ada dan sesuaikan dengan konfigurasi lokal Anda, terutama untuk koneksi database.

    ```sh
    cp .env.example .env
    ```

3.  **Isi file `.env` dengan kredensial Anda:**

    ```env
    # Server Configuration
    SERVER_PORT=8080

    # Database Configuration
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=password_anda
    DB_NAME=nama_database_anda
    ```

4.  **Unduh dependensi proyek:**
    Perintah ini akan mengunduh semua modul yang dibutuhkan yang terdaftar di `go.mod`.

    ```sh
    go mod tidy
    ```

5.  **Jalankan aplikasi:**
    ```sh
    go run cmd/main.go
    ```
    Server sekarang akan berjalan di `http://localhost:8080` (atau port yang Anda tentukan di `.env`).

## 🧪 Pengujian API

Untuk menguji endpoint API, Anda dapat menggunakan koleksi Postman yang telah disediakan.

1.  Buka aplikasi Postman Anda.
2.  Klik **Import** -> **File** -> **Upload Files**.
3.  Pilih file `evermos virtual intern.postman_collection.json` dari direktori proyek.
4.  Setelah terimpor, Anda akan melihat daftar endpoint API yang siap untuk diuji.
