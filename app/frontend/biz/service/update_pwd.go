package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

type UpdatePwdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePwdService(Context context.Context, RequestContext *app.RequestContext) *UpdatePwdService {
	return &UpdatePwdService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePwdService) Run(req *auth.UpdatePwdReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
