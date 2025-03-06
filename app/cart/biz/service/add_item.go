package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/meilingluolingluo/gomall/app/cart/biz/model"
	"github.com/meilingluolingluo/gomall/app/cart/rpc"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	cart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
		Id: req.Item.ProductId,
	})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(5003002, "product id is required")
	}
	cartItems := model.Cart{
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
		UserId:    req.UserId,
	}

	if err := model.AddCart(mysql.DB, s.ctx, &cartItems); err != nil {
		return nil, kerrors.NewBizStatusError(5003003, err.Error())
	}
	resp = &cart.AddItemResp{}

	return resp, nil
}
