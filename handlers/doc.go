// Package handlers berisi handler HTTP Gin untuk semua endpoint gocommerce:
// produk, kategori produk, transaksi, login, dan register, serta pembuatan token JWT.
//
// Setiap fungsi handler berbentuk func(db *gorm.DB) gin.HandlerFunc sehingga
// koneksi database dikirim lewat parameter dan mudah diganti saat pengujian.
package handlers
