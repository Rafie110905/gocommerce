package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB adalah koneksi database global yang diisi oleh InitDB.
// Nilainya nil sampai InitDB berhasil dipanggil.
var DB *gorm.DB

// InitDB membuka koneksi ke database MySQL "ecomm" di 127.0.0.1:3306,
// menyimpannya di variabel global DB, lalu mengembalikannya.
//
// Jika koneksi gagal, InitDB mengembalikan (nil, error) dan DB tidak diubah.
// Pemanggil bertanggung jawab menutup koneksi dengan db.Close().
//
// Catatan: DSN (termasuk username dan password) saat ini ditulis langsung di
// kode. Untuk produksi, pindahkan ke environment variable.
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Password@tcp(127.0.0.1:3306)/ecomm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}
