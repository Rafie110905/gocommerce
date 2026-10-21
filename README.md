# Gocommerce

REST API e-commerce sederhana yang ditulis dengan Go. Mendukung manajemen produk, kategori produk, transaksi beserta item-nya, serta registrasi dan login dengan JWT.

**Teknologi:** Go 1.21, [Gin](https://github.com/gin-gonic/gin) (HTTP router), [GORM v1](https://github.com/jinzhu/gorm) (ORM), MySQL, [jwt-go](https://github.com/dgrijalva/jwt-go) (autentikasi), [testify](https://github.com/stretchr/testify) (test).

## Daftar isi

1. [Struktur proyek](#struktur-proyek)
2. [Prasyarat](#prasyarat)
3. [Menjalankan aplikasi](#menjalankan-aplikasi)
4. [Dokumentasi API](#dokumentasi-api)
5. [Autentikasi](#autentikasi)
6. [Menjalankan test](#menjalankan-test)
7. [Dokumentasi kode (GoDoc)](#dokumentasi-kode-godoc)
8. [Keterbatasan yang diketahui](#keterbatasan-yang-diketahui)
9. [Perintah berguna](#perintah-berguna)

## Struktur proyek

```
gocommerce/
├── main.go                  # entry point, mendaftarkan rute, port 5000
├── configs/config.go        # koneksi database MySQL
├── middlewares/             # AuthMiddleware (validasi JWT)
├── handlers/                # handler HTTP: produk, kategori, transaksi, login, register
├── models/                  # struct data (juga bentuk JSON request/response)
├── migrations/              # AutoMigrate tabel
├── seeders/                 # data contoh
├── docs/
│   └── gocommerce.postman_collection.json
├── tempCheck.go             # skrip debugging sementara (lihat Keterbatasan)
└── tempCheck2.go            # skrip debugging sementara
```

## Prasyarat

- Go 1.21 atau lebih baru
- MySQL berjalan di `127.0.0.1:3306`
- Database kosong bernama `ecomm`

Koneksi database ditulis di `configs/config.go` dengan user `root` dan password `Password`. Sesuaikan dengan MySQL Anda sebelum menjalankan.

## Menjalankan aplikasi

```bash
# 1. Buat database
mysql -u root -p -e "CREATE DATABASE ecomm CHARACTER SET utf8mb4;"

# 2. Unduh dependensi
go mod download

# 3. Jalankan server
go run main.go
```

Server berjalan di `http://localhost:5000`. Tabel dibuat otomatis saat startup (`migrations.Migrate`).

Untuk mengisi data contoh (2 kategori, 2 user, 2 produk, 2 transaksi), hapus tanda komentar pada baris `// seeders.Seed(db)` di `main.go` **satu kali saja**. Seeder tidak mengecek data yang sudah ada, jadi menjalankannya berulang akan menggandakan data. Akun contoh: `user1` / `password1` dan `user2` / `password2`.

> Gunakan `go run main.go`, bukan `go run .`. Lihat bagian [Keterbatasan](#keterbatasan-yang-diketahui).

## Dokumentasi API

Base URL: `http://localhost:5000`. Semua request dan response berformat JSON.

| Method | Endpoint | Auth | Keterangan |
|---|---|---|---|
| POST | `/register` | Tidak | Daftar user baru, mengembalikan token |
| POST | `/login` | Tidak | Login, mengembalikan token |
| GET | `/products` | **Ya** | Daftar semua produk |
| GET | `/products/:id` | Tidak | Detail satu produk |
| POST | `/products` | Tidak | Buat produk |
| PUT | `/products/:id` | Tidak | Ubah produk |
| DELETE | `/products/:id` | Tidak | Hapus produk |
| GET | `/product-categories` | Tidak | Daftar semua kategori |
| GET | `/product-categories/:id` | Tidak | Detail satu kategori |
| POST | `/product-categories` | Tidak | Buat kategori |
| PUT | `/product-categories/:id` | Tidak | Ubah kategori |
| DELETE | `/product-categories/:id` | Tidak | Hapus kategori |
| POST | `/transactions` | Tidak | Buat transaksi beserta item |
| GET | `/transactions/:id` | Tidak | Detail transaksi beserta item |
| GET | `/debug/pprof/*` | Tidak | Profiling performa (pprof) |

### Contoh request dan response

**Register**
```http
POST /register
Content-Type: application/json

{"username": "budi", "email": "budi@example.com", "password": "rahasia123"}
```
```json
{"message": "User registered successfully", "token": "eyJhbGciOiJIUzI1NiIs..."}
```

**Login**
```http
POST /login
Content-Type: application/json

{"username": "user1", "password": "password1"}
```
```json
{"token": "eyJhbGciOiJIUzI1NiIs..."}
```

**Daftar produk (butuh token)**
```http
GET /products
Authorization: eyJhbGciOiJIUzI1NiIs...
```
```json
[{"id": 1, "name": "Product 1", "category_id": 1, "price": 0}]
```

**Buat produk**
```http
POST /products
Content-Type: application/json

{"name": "Kaos Polos", "category_id": 1, "price": 75000}
```
```json
{"id": 3, "name": "Kaos Polos", "category_id": 1, "price": 75000}
```

**Buat transaksi**
```http
POST /transactions
Content-Type: application/json

{
  "user_id": 1,
  "amount": 150000,
  "items": [{"product_id": 1, "quantity": 2, "price": 75000}]
}
```
```json
{
  "id": 3, "user_id": 1, "amount": 150000, "price": 0,
  "items": [{"id": 3, "transaction_id": 3, "product_id": 1, "quantity": 2, "price": 75000}]
}
```

### Kode status

| Kode | Arti |
|---|---|
| 200 | Berhasil |
| 201 | Data berhasil dibuat |
| 400 | Body tidak valid, username sudah ada, name kategori kosong, atau `product_id` pada transaksi tidak ditemukan |
| 401 | Token tidak ada/tidak valid, atau username login tidak ditemukan |
| 404 | Data dengan `id` tersebut tidak ditemukan |
| 500 | Kesalahan server |

Semua error berbentuk `{"message": "..."}`.

### Koleksi Postman

Import `docs/gocommerce.postman_collection.json` ke Postman (**Import** lalu pilih file). Koleksi sudah berisi:

- variabel `baseUrl` (default `http://localhost:5000`) dan `token`,
- seluruh 14 endpoint aplikasi dengan deskripsi, contoh body, dan contoh response,
- skrip test yang otomatis menyimpan token ke variabel `token` setelah **Register** atau **Login** berhasil,
- test status kode pada setiap request.

Urutan yang disarankan: **Login**, lalu **List Products**.

## Autentikasi

API memakai JWT (HS256). Alurnya:

1. Panggil `POST /login` atau `POST /register` untuk mendapat `token`.
2. Kirim token pada header `Authorization` untuk rute yang dilindungi.

**Penting:** isi header adalah token mentah, **tanpa** awalan `Bearer `.

```
Authorization: eyJhbGciOiJIUzI1NiIs...
```

Saat ini hanya `GET /products` yang dilindungi.

## Menjalankan test

```bash
go test ./handlers/ -v
```

Package `handlers` menyiapkan database lewat `init()`: mencoba MySQL lebih dulu, lalu fallback ke SQLite in-memory beserta migrasi dan data seeder. Driver SQLite (`go-sqlite3`) memerlukan CGO dan compiler C (misalnya gcc).

## Dokumentasi kode (GoDoc)

Seluruh package, fungsi, dan struct sudah diberi komentar dokumentasi. Untuk membacanya di terminal:

```bash
go doc ./handlers
go doc ./handlers Login
go doc ./models User
```

Atau jalankan server dokumentasi lokal:

```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .
```

Lalu buka `http://localhost:8080`. Jika repo ini diunggah ke GitHub, dokumentasi publik otomatis tersedia di `https://pkg.go.dev/github.com/<username>/<repo>`.

## Keterbatasan yang diketahui

Proyek ini adalah bahan belajar. Hal-hal berikut perlu diperbaiki sebelum dipakai di produksi:

**Keamanan**
- Password disimpan sebagai teks biasa dan **tidak diperiksa saat login**. Siapa pun yang tahu username bisa mendapat token. Gunakan `bcrypt` untuk hash dan verifikasi.
- Kunci JWT (`your-secret-key`) dan kredensial database ditulis langsung di kode, dan kunci JWT terduplikasi di dua package. Pindahkan ke environment variable.
- Token tidak punya waktu kedaluwarsa.
- Hanya `GET /products` yang dilindungi. Rute tulis (POST, PUT, DELETE) dan transaksi terbuka untuk siapa saja.

**Perilaku**
- `amount` transaksi diambil dari body klien, tidak dihitung dari item.
- Error dari `db.Create` pada produk, kategori, dan transaksi belum diperiksa.
- `category_id` pada produk tidak divalidasi terhadap tabel kategori.
- Endpoint daftar mengembalikan `null` (bukan `[]`) saat tabel kosong.
- Data seeder untuk produk belum memiliki harga (`price` bernilai 0).

**Struktur repo**
- `main.go`, `tempCheck.go`, dan `tempCheck2.go` sama-sama `package main` dan masing-masing punya `func main()`. Akibatnya `go build` dan `go run .` gagal dengan error `main redeclared`. Gunakan `go run main.go`, atau pindahkan/hapus kedua file `tempCheck`.
- `gocommerce.exe` adalah binary hasil build dan sebaiknya tidak disimpan di repo (tambahkan ke `.gitignore`).
- `documentation.html` masih kosong.
- `package handlers` membuka koneksi database di `init()`, sehingga ikut berjalan saat aplikasi utama dimulai.

## Perintah berguna

```bash
# Profiling CPU lewat pprof (server harus berjalan)
go tool pprof http://localhost:6060/debug/pprof/profile

# Menambahkan binary Go ke PATH
export PATH=$(go env GOPATH)/bin:$PATH
```

> Catatan: rute pprof pada `main.go` dilayani lewat port aplikasi (5000), sehingga alamat profil yang sesuai adalah `http://localhost:5000/debug/pprof/profile`.
