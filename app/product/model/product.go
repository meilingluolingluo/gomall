package model

type Product struct {
	Base
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Picture     string `json:"picture" gorm:"not null"`
	Price       int32  `json:"price" gorm:"not null"`

	Categories []Category `json:"categories" gorm:"many2many:product_categories;"`
}

func (p Product) TableName() string {
	return "products"
}
