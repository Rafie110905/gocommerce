package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ListProductCategories menangani GET /product-categories dan mengembalikan
// seluruh kategori produk sebagai array JSON (200). Bernilai null jika belum ada data.
func ListProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.ProductCategory
		db.Find(&categories)
		c.JSON(200, categories)
	}
}

// GetProductCategory menangani GET /product-categories/:id.
//
// Response: 200 dengan satu kategori, atau 404 {"message": "Category not found"}.
func GetProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}
		c.JSON(200, category)
	}
}

// CreateProductCategory menangani POST /product-categories.
//
// Request body (JSON):
//
// 	{"name": "Pakaian"}
//
// Response: 201 dengan kategori yang dibuat. Mengembalikan 400 jika body tidak
// valid ("Invalid input") atau name kosong ("Name is required").
func CreateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.ProductCategory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		if input.Name == "" {
			c.JSON(400, gin.H{"message": "Name is required"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

// UpdateProductCategory menangani PUT /product-categories/:id.
//
// Response: 200 dengan kategori yang diperbarui, 404 jika id tidak ada,
// 400 jika body tidak valid.
func UpdateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}

		var input models.ProductCategory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&category).Updates(input)
		c.JSON(200, category)
	}
}

// DeleteProductCategory menangani DELETE /product-categories/:id.
//
// Response: 200 {"message": "Category deleted"}, atau 404 jika id tidak ada.
func DeleteProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}

		db.Delete(&category)
		c.JSON(200, gin.H{"message": "Category deleted"})
	}
}
