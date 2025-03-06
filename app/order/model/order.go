package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(255);uniqueIndex"`
	UserId       uint        `gorm:"type:int"`
	UserCurrency string      `gorm:"type:varchar(255)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrders(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	if err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 删除订单函数
func DeleteOrder(ctx context.Context, db *gorm.DB, userId uint, orderId string) error {
	result := db.WithContext(ctx).Where("user_id = ? AND order_id = ?", userId, orderId).Delete(&Order{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// 更新订单函数
func UpdateOrder(ctx context.Context, db *gorm.DB, userId uint, orderId string, consignee *Consignee) error {
	result := db.WithContext(ctx).Model(&Order{}).Where("user_id = ? AND order_id = ?", userId, orderId).Updates(Order{
		Consignee: *consignee,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
