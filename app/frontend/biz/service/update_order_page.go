package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
)

type UpdateOrderPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderPageService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderPageService {
	return &UpdateOrderPageService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderPageService) Run(req *order.UpdateOrderPageReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	return utils.H{
		"OrderId": req.OrderId,
		"UserId":  userId,
	}, nil
}
