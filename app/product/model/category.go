package model

type Category struct {
	Base
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`

	Products []Product `json:"products" gorm:"many2many:product_categories;"`
}
