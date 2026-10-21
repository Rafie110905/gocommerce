package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Login menangani POST /login.
//
// Request body (JSON):
//
// 	{"username": "user1", "password": "password1"}
//
// Response:
//   - 200: {"token": "<jwt>"} jika username ditemukan.
//   - 400: {"message": "Invalid input"} jika body bukan JSON yang valid.
//   - 401: {"message": "Invalid credentials"} jika username tidak ditemukan.
//   - 500: {"message": "Internal Server Error"} jika pembuatan token gagal.
//
// Catatan: saat ini password belum diperiksa. Siapa pun yang tahu username
// yang valid bisa mendapatkan token.
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"message": "Invalid credentials"})
			return
		}

		// Pastikan untuk memeriksa kata sandi yang benar di sini

		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}
