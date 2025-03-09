package service

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	"github.com/meilingluolingluo/gomall/app/frontend/types"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: userId,
	})
	log.Printf("userId: %d", userId)
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		log.Printf("order: %v", v)
		for _, v := range v.OrderItems {
			total += v.Cost
			i := v.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			log.Printf("productResp: %v", productResp)
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product
			items = append(items, types.OrderItem{
				Id:          int32(i.ProductId),
				ProductName: p.Name,
				Quantity:    int32(i.Quantity),
				Cost:        float64(v.Cost),
				Picture:     p.Picture,
			})
		}
		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:    v.OrderId,
			CreateData: created.Format("2006-01-02 15:04:05"),
			OrderItems: items,
			Cost:       float64(total),
			Consignee: types.Consignee{
				City:          v.Address.City,
				Country:       v.Address.Country,
				Email:         v.Email,
				State:         v.Address.State,
				StreetAddress: v.Address.StreetAddress,
				ZipCode:       int32(v.Address.ZipCode),
			},
			PaymentStatus: v.Paymentstatus,
		})
	}

	log.Printf("orders list %v", list)
	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
