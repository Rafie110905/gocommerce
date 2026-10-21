package models

// Transaction adalah satu transaksi pembelian oleh seorang User. Disimpan di tabel "transactions".
type Transaction struct {
	// ID adalah primary key, dibuat otomatis oleh database.
	ID     uint              `gorm:"primary_key" json:"id"`
	// UserID adalah ID User yang melakukan transaksi.
	UserID uint              `json:"user_id"`
	// Amount adalah total nilai transaksi. Diisi oleh klien, tidak dihitung otomatis.
	Amount float64           `json:"amount"`
	// Items adalah daftar TransactionItem milik transaksi ini (relasi lewat TransactionID).
	Items  []TransactionItem `gorm:"foreignkey:TransactionID" json:"items"`
	// Price adalah atribut harga tambahan pada transaksi.
	Price  float64           `json:"price"` // Tambahkan atribut Price di sini
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
