package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mallowww/tuan-redis/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	products, err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)
}

func initDB() *gorm.DB {
	dsn := "root:notApengu1ndb@tcp(localhost:3306)/penguin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
