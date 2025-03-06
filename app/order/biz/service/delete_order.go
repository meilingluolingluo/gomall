package service

import (
	"context"
	"errors"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/order/model"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	// 校验请求参数
	if req.UserId == 0 || req.OrderId == "" {
		return nil, errors.New("invalid user ID or order ID")
	}

	// 使用事务执行删除操作
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 调用删除逻辑
		if err := model.DeleteOrder(s.ctx, tx, uint(req.UserId), req.OrderId); err != nil {
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
	resp = &order.DeleteOrderResp{
		Order: &order.OrderResult{
			OrderId: req.OrderId,
		},
	}
	return
}
