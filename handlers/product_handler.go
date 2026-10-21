package handlers

import (
	"gocommerce/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// contoh lama
// func ListProducts(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var products []models.Product
// 		db.Find(&products)
// 		c.JSON(200, products)
// 	}
// }

// contoh baru dengan goroutine

// ListProducts menangani GET /products dan mengembalikan seluruh produk.
// Rute ini dilindungi AuthMiddleware, jadi request harus menyertakan header
// "Authorization" berisi token JWT.
//
// Response: 200 dengan array produk (bernilai null jika belum ada data).
// Query ke database dijalankan di goroutine dan ditunggu dengan sync.WaitGroup.
//
// Anotasi swag untuk dokumentasi Swagger:
// @Summary Get a list of products
// @Description Retrieve a list of products from the database.
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		// Menambahkan satu goroutine ke WaitGroup
		wg.Add(1)

		// Memulai goroutine untuk melakukan operasi yang membutuhkan waktu lama
		go func() {
			defer wg.Done() // Menandai bahwa goroutine telah selesai
			db.Find(&products)
		}()

		// Menunggu goroutine selesai
		wg.Wait()

		c.JSON(200, products)
	}
}

// GetProduct menangani GET /products/:id.
//
// Response: 200 dengan satu produk, atau 404 {"message": "Product not found"}
// jika id tidak ada.
func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(200, product)
	}
}

// CreateProduct menangani POST /products.
//
// Request body (JSON):
//
// 	{"name": "Kaos Polos", "category_id": 1, "price": 75000}
//
// Response: 201 dengan produk yang dibuat (termasuk id), atau 400
// {"message": "Invalid input"} jika body tidak valid. Hasil db.Create belum
// diperiksa untuk error, dan category_id tidak divalidasi terhadap tabel kategori.
func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

// UpdateProduct menangani PUT /products/:id.
//
// Body berisi field yang ingin diubah. Field yang bernilai nol (misalnya "" atau
// 0) diabaikan oleh GORM Updates sehingga tidak bisa dipakai untuk mengosongkan nilai.
//
// Response: 200 dengan produk, 404 jika id tidak ada, 400 jika body tidak valid.
// Urutan pengecekan: id dicari dulu, baru body divalidasi.
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&product).Updates(input)
		c.JSON(200, product)
	}
}

// DeleteProduct menangani DELETE /products/:id.
//
// Response: 200 {"message": "Product deleted"}, atau 404 jika id tidak ada.
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		db.Delete(&product)
		c.JSON(200, gin.H{"message": "Product deleted"})
	}
}
