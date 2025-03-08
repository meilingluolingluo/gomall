package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	// var resp = make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture001.png"},
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture006.png"},
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture009.png"},
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture012.png"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["items"] = items
	// return resp, nil

	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})

	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Category",
		"items": p.Products,
	}, nil
}
