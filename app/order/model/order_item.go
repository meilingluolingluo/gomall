package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ProductId    uint    `gorm:"type:int(11)"`
	OrderIdRefer string  `gorm:"type:varchar(255);index"`
	Quantity     uint    `gorm:"type:int(11)"`
	Cost         float64 `gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
