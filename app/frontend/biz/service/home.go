package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {

	// 添加错误处理
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		klog.Errorf("Failed to list products: %v", err)
		return nil, err
	}
	return utils.H{
		"Title": "Category",
		"items": p.Products,
	}, nil
}
