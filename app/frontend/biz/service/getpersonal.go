package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

type GetpersonalService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetpersonalService(Context context.Context, RequestContext *app.RequestContext) *GetpersonalService {
	return &GetpersonalService{RequestContext: RequestContext, Context: Context}
}

func (h *GetpersonalService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
