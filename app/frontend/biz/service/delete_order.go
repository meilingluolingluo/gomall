package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
)

type DeleteOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteOrderService(Context context.Context, RequestContext *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
