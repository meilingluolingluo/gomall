package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	rpccart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 调用RPC服务
	if _, err = rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{
		UserId: uint32(h.Context.Value(frontendutils.SessionUserId).(float64)),
	}); err != nil {
		return nil, fmt.Errorf("清空购物车失败: %w", err)
	}

	// 返回模板数据
	return utils.H{
		"message": "购物车已清空",
	}, nil
}
