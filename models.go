package main

import "gorm.io/gorm"

type Product struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}

type Category struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

type Cart struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Products []Product `json:"products" gorm:"many2many:cart_products;"`
}

func ScopeByCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}

func ScopeExpensiveProducts(minPrice float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price >= ?", minPrice)
	}
}
