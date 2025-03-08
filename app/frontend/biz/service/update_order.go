package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
