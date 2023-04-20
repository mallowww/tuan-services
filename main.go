package main

import (
	"fmt"

	"github.com/mallowww/tuan-redis/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	productRepo := repositories.NewProductRepositoryDB(db)
	products, err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)
}

func initDB() *gorm.DB {
	// dial := mysql.Open("root:notApengu1ndb@tcp(localhost:3350)/penguin")
	// db, err := gorm.Open(dial, &gorm.Config{})
	dsn := "root:notApengu1ndb@tcp(localhost:3306)/penguin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
