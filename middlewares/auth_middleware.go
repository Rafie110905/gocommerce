// Package middlewares berisi middleware Gin untuk aplikasi gocommerce.
//
// Saat ini hanya ada satu middleware, yaitu AuthMiddleware untuk memvalidasi
// token JWT pada rute yang dilindungi.
package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwtKey adalah kunci rahasia untuk memverifikasi tanda tangan token JWT (HS256).
// Nilainya harus sama dengan jwtKey di package handlers, yang dipakai untuk
// membuat token.
var jwtKey = []byte("your-secret-key") // Ganti dengan kunci rahasia Anda

// AuthMiddleware mengembalikan middleware Gin yang melindungi sebuah rute dengan JWT.
//
// Cara kerja:
//   1. Membaca header "Authorization". Isinya adalah token mentah, tanpa
//      awalan "Bearer ".
//   2. Memverifikasi token dengan jwtKey.
//   3. Jika valid, menyimpan klaim "user_id" (bertipe uint) ke context
//      sehingga handler berikutnya bisa membacanya dengan c.Get("user_id").
//
// Jika header kosong atau token tidak valid, middleware menghentikan request
// dengan status 401 dan body {"message": "Unauthorized"}.
//
// Contoh pemakaian:
//
// 	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Jika token valid, Anda dapat mengekstrak informasi pengguna dari token di sini
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()
	}
}
