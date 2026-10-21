package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetTransactionWithItems menangani GET /transactions/:id.
//
// Mengembalikan transaksi beserta seluruh item-nya (relasi Items dimuat dengan
// Preload). Response: 200 dengan transaksi, atau 404
// {"message": "Transaction not found"} jika id tidak ada.
func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, transaction)
	}
}

// CreateTransaction menangani POST /transactions.
//
// Request body (JSON):
//
// 	{
// 	  "user_id": 1,
// 	  "amount": 150000,
// 	  "items": [{"product_id": 1, "quantity": 2, "price": 75000}]
// 	}
//
// Setiap product_id pada items diperiksa keberadaannya sebelum transaksi disimpan.
//
// Response: 201 dengan transaksi yang dibuat (termasuk id), 400
// {"message": "Invalid input"} jika body tidak valid, atau 400
// {"message": "Invalid product ID"} jika ada produk yang tidak ditemukan.
//
// Catatan: amount diambil dari body apa adanya, tidak dihitung ulang dari items.
func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Transaction
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		// Jika Anda ingin mengaitkan item transaksi, pastikan item-item tersebut ada dalam database
		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(400, gin.H{"message": "Invalid product ID"})
				return
			}
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}
