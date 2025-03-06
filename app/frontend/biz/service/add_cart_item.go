package service

import (
	"context"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	rpccart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	return
}
