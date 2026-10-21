// Package migrations berisi pembuatan dan pembaruan skema tabel database.
package migrations

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

// Migrate membuat atau memperbarui tabel products, product_categories,
// transactions, transaction_items, dan users memakai GORM AutoMigrate.
//
// AutoMigrate hanya menambah tabel dan kolom baru. Fungsi ini tidak pernah
// menghapus kolom atau data yang sudah ada, sehingga aman dipanggil setiap
// aplikasi dimulai.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.User{},
	)
}
