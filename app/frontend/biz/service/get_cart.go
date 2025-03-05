package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	rpccart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
	"strconv"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(h.Context.Value(frontendutils.SessionUserId).(float64)),
	})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.GetProductId()})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "购物车",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
