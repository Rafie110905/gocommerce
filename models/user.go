package models

// User adalah akun pengguna yang bisa login ke API. Disimpan di tabel "users".
type User struct {
	// ID adalah primary key, dibuat otomatis oleh database.
	ID       uint   `gorm:"primary_key" json:"id"`
	// Username adalah nama login. Dipakai untuk mencari user saat login dan register.
	Username string `json:"username"`
	// Email adalah alamat email pengguna.
	Email    string `json:"email"`
	// Password adalah kata sandi pengguna. Saat ini disimpan sebagai teks biasa (belum di-hash).
	Password string `json:"password"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
