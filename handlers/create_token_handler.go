package handlers

import "github.com/dgrijalva/jwt-go"

// jwtKey adalah kunci rahasia untuk menandatangani token JWT (HS256).
// Harus sama dengan jwtKey di package middlewares agar token bisa diverifikasi.
var jwtKey = []byte("your-secret-key") // Ganti dengan kunci rahasia Anda

// CreateToken membuat token JWT yang ditandatangani dengan HS256 untuk userID.
//
// Token hanya memuat satu klaim, yaitu "user_id". Token tidak punya waktu
// kedaluwarsa (klaim "exp"). Mengembalikan error jika proses penandatanganan gagal.
func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
