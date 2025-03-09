package service

import (
	"context"
	"errors"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/order/model"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// Finish your business logic.
	// 校验请求参数
	if req.UserId == 0 || req.OrderId == "" {
		return nil, errors.New("invalid user ID or order ID")
	}

	// 使用事务执行更新操作
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 填充 Consignee 结构体
		consignee := &model.Consignee{
			Email:         req.Email,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		}

		// 调用更新逻辑
		if err := model.UpdateOrder(s.ctx, tx, uint(req.UserId), req.OrderId, consignee); err != nil {
			return err
		}
		return nil
	})

	// 处理错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	// 返回响应
	resp = &order.UpdateOrderResp{
		Order: &order.OrderResult{
			OrderId: req.OrderId,
		},
	}
	return
}
