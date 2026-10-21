package models

// ProductCategory adalah kategori untuk mengelompokkan produk. Disimpan di tabel "product_categories".
type ProductCategory struct {
	// ID adalah primary key, dibuat otomatis oleh database.
	ID   uint   `gorm:"primary_key" json:"id"`
	// Name adalah nama kategori. Wajib diisi saat membuat kategori lewat API.
	Name string `json:"name"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
