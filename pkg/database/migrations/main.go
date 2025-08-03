package main

import (
	"Evermos-Virtual-Intern/internal/entity" // Ganti dengan module name kamu
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_evermos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Alamat{},
		&entity.Toko{},
		&entity.Category{},
		&entity.Produk{},
		&entity.FotoProduk{},
		&entity.LogProduk{},
		&entity.Trx{},
		&entity.DetailTrx{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	fmt.Println("Database migration success!")
}
