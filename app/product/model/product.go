package model

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Picture     string     `json:"picture" gorm:"not null"`
	Price       float32    `json:"price" gorm:"not null"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category;"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetByID(productid uint32) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&product, productid).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	log.Printf("q = %s", q)
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("name LIKE ? OR description LIKE ?", "%"+q+"%", "%"+q+"%").Find(&products).Error
	log.Printf("products = %v", products)
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}
