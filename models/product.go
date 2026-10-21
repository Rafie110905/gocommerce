package models

// Product adalah produk yang dijual di toko. Disimpan di tabel "products".
type Product struct {
	// ID adalah primary key, dibuat otomatis oleh database.
	ID         uint    `gorm:"primary_key" json:"id"`
	// Name adalah nama produk.
	Name       string  `json:"name"`
	// CategoryID adalah ID ProductCategory tempat produk ini berada.
	CategoryID uint    `json:"category_id"`
	// Price adalah harga satuan produk.
	Price      float64 `json:"price"` // Tambahkan atribut Price di sini
}
