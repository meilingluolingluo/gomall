package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/meilingluolingluo/gomall/app/cart/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/cart/biz/model"
	cart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// 获取购物车列表
	carts, err := model.GetCartByUserId(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range carts {
		items = append(items, &cart.CartItem{ProductId: v.ProductId, Quantity: v.Qty})
	}
	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.GetUserId(), Items: items}}, nil

}
