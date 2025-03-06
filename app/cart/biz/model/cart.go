package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       int32  `json:"qty"`
}

func (c Cart) TableName() string {
	return "cart"
}

func GetCartByUserId(db *gorm.DB, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}

/*
	func AddCart(db *gorm.DB, ctx context.Context, c *Cart) error {
		var find Cart
		// Log the operation for debugging
		klog.Debugf("Adding/updating cart item: user_id=%d, product_id=%d, quantity=%d", c.UserId, c.ProductId, c.Qty)
		err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&find).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if find.ID != 0 {
			err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
		} else {
			err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
		}
		return err
	}
*/
func AddCart(db *gorm.DB, ctx context.Context, c *Cart) error {
	if c == nil {
		return fmt.Errorf("cart item cannot be nil")
	}

	// 输入参数验证
	if c.UserId <= 0 {
		return fmt.Errorf("invalid user ID: %d", c.UserId)
	}
	if c.ProductId <= 0 {
		return fmt.Errorf("invalid product ID: %d", c.ProductId)
	}
	if c.Qty <= 0 {
		return fmt.Errorf("quantity must be positive: %d", c.Qty)
	}

	// Log
	klog.Debugf("Adding/updating cart item: user_id=%d, product_id=%d, quantity=%d", c.UserId, c.ProductId, c.Qty)

	// transaction
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingCart Cart
		// Use a single query with index optimization
		err := tx.Model(&Cart{}).
			Where("user_id = ? AND product_id = ?", c.UserId, c.ProductId).
			First(&existingCart).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Errorf("Failed to query cart item: %v", err)
			return fmt.Errorf("failed to query cart item: %v", err)
		}

		if existingCart.ID != 0 {
			// Update existing cart item
			err = tx.Model(&Cart{}).
				Where("user_id = ? AND product_id = ?", c.UserId, c.ProductId).
				UpdateColumn("qty", gorm.Expr("qty + ?", c.Qty)).
				UpdateColumn("updated_at", time.Now()).Error
			if err != nil {
				klog.Errorf("Failed to update cart item: %v", err)
				return fmt.Errorf("failed to update cart item: %v", err)
			}
		} else {
			// Create new cart item
			c.CreatedAt = time.Now()
			c.UpdatedAt = time.Now()
			err = tx.Create(c).Error
			if err != nil {
				klog.Errorf("Failed to create cart item: %v", err)
				return fmt.Errorf("failed to create cart item: %v", err)
			}
		}

		return nil
	})
}
func EmptyCart(db *gorm.DB, ctx context.Context, userId uint32) error {
	if userId == 0 {
		return errors.New("user_id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}
