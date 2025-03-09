package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendutils.GetUserIdFromCtx(ctx)
	username := frontendutils.GetUserNameFromCtx(ctx)
	content["user_id"] = userId
	content["username"] = username
	if userId > 0 {

		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: frontendutils.GetUserIdFromCtx(ctx),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}

	return content
}
