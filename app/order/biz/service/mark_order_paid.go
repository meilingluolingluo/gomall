package service

import (
	"context"
	"errors"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/order/model"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	// 校验请求参数
	if req.UserId == 0 || req.OrderId == "" {
		return nil, errors.New("invalid user ID or order ID")
	}

	// 使用事务执行更新操作
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {

		// 调用更新逻辑
		if err := model.MarkOrderPaid(s.ctx, tx, uint(req.UserId), req.OrderId); err != nil {
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
	resp = &order.MarkOrderPaidResp{}
	return
}
