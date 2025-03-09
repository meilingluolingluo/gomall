package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	deleteorder "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

type DeleteOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteOrderService(Context context.Context, RequestContext *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	// log.Printf("userId: %d", userId)
	// log.Printf("OrderId: %s", req.OrderId)
	deleteResp, err := rpc.OrderClient.DeleteOrder(h.Context, &deleteorder.DeleteOrderReq{
		OrderId: req.OrderId,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"message": fmt.Sprintf("Order %s has been deleted", deleteResp.Order.OrderId),
	}, nil

}
