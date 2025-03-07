package model

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
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

// CreateOrderItem 创建订单商品
func CreateOrderItem(ctx context.Context, tx *gorm.DB, items *[]OrderItem) error {
	result := tx.Create(items)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return kerrors.NewGRPCBizStatusError(500002, "no rows affected while creating order items")
	}
	return nil
}

// DeleteOrderItem 删除指定订单ID下的所有订单商品
func DeleteOrderItem(ctx context.Context, tx *gorm.DB, orderId string) error {
	result := tx.Where("order_id_refer = ?", orderId).Delete(&OrderItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return kerrors.NewGRPCBizStatusError(500003, "no rows affected while deleting order items")
	}
	return nil
}
