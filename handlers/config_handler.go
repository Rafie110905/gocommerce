package handlers

import (
	"fmt"
	"gocommerce/migrations"
	"gocommerce/seeders"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// db adalah koneksi database tingkat package yang hanya dipakai oleh test
// (lihat transaction_handler_test.go). Handler produksi menerima koneksi
// lewat parameter, bukan lewat variabel ini.
var db *gorm.DB

// init menyiapkan database untuk lingkungan test.
//
// Urutannya: coba MySQL terlebih dahulu. Jika tidak terjangkau, fallback ke
// SQLite in-memory dengan satu koneksi agar data tetap ada selama proses
// berjalan. Setelah itu tabel dimigrasi dan data contoh diisi (seeders.Seed).
//
// Catatan: karena berupa init(), kode ini berjalan setiap kali package handlers
// di-import, termasuk saat aplikasi utama berjalan.
func init() {
	var err error

	// Try primary MySQL database first.
	db, err = gorm.Open("mysql", "root:Password@tcp(127.0.0.1:3306)/ecomm?charset=utf8mb4&parseTime=True&loc=Local")

	// Check connectivity – db.Ping reveals whether the database is actually reachable.
	if err == nil {
		if pingErr := db.DB().Ping(); pingErr != nil {
			err = pingErr
		}
	}

	if err != nil {
		// Fallback to an in-memory SQLite database for local/testing environments.
		db, err = gorm.Open("sqlite3", ":memory:")
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
		// Keep a single connection so the in-memory database persists for the life of the process.
		db.DB().SetMaxOpenConns(1)
	}

	// Run migrations and seed data so that tests have the required tables & records.
	fmt.Println("[handlers] running migrations and seeding test data")
	migrations.Migrate(db)
	seeders.Seed(db)
}
