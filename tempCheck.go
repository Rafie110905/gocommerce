package main

import (
    "fmt"
    "gocommerce/migrations"
    "gocommerce/seeders"
    "gocommerce/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
    db, err := gorm.Open("sqlite3", ":memory:")
    if err != nil {
        panic(err)
    }
    migrations.Migrate(db)
    seeders.Seed(db)

    var tx models.Transaction
    if err := db.Preload("Items").First(&tx, 1).Error; err != nil {
        fmt.Println("query error:", err)
    } else {
        fmt.Printf("Transaction found: %+v\n", tx)
    }
}