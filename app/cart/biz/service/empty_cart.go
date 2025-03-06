package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/meilingluolingluo/gomall/app/cart/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/cart/biz/model"
	cart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyCart(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5003001, "product id is required")
	}
	return &cart.EmptyCartResp{}, nil
}
