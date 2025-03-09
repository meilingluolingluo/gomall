package cart

import (
	"context"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/service"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/utils"
	cart "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}
	c.Redirect(consts.StatusFound, []byte("/cart"))
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

// EmptyCart .
// @router /cart/empty [POST]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var req common.Empty
	if err := c.BindAndValidate(&req); err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	_, err := service.NewEmptyCartService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.Redirect(consts.StatusFound, []byte("/cart"))
}
