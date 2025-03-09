package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	updateorder "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	updateResp, err := rpc.OrderClient.UpdateOrder(h.Context, &updateorder.UpdateOrderReq{
		OrderId: req.OrderId,
		UserId:  userId,
		Email:   req.Email,
		Address: &updateorder.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"message": fmt.Sprintf("Order %s has been updated", updateResp.Order.OrderId),
	}, nil
}
