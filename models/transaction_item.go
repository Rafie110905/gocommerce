package models

// TransactionItem adalah satu baris produk di dalam Transaction. Disimpan di tabel "transaction_items".
type TransactionItem struct {
	// ID adalah primary key, dibuat otomatis oleh database.
	ID            uint    `gorm:"primary_key" json:"id"`
	// TransactionID adalah ID Transaction pemilik item ini.
	TransactionID uint    `json:"transaction_id"`
	// ProductID adalah ID Product yang dibeli. Harus sudah ada di tabel products.
	ProductID     uint    `json:"product_id"`
	// Quantity adalah jumlah unit produk yang dibeli.
	Quantity      uint    `json:"quantity"`
	// Price adalah harga satuan produk saat transaksi terjadi.
	Price         float64 `json:"price"` // Tambahkan atribut Price di sini
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
