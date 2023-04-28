package repositories

import (
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) GetProduct() (products []product, err error) {
	err = r.db.Order("Quantity DESC").Limit(20).Find(&products).Error
	return products, err
}
