package product

import (
	"context"

	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/service"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/utils"
	product "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/product"
)

// GetProduct .
// @router /product [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductReq

	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, resp))
}

// SearchProducs .
// @router /search [GET]
func SearchProducs(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	resp, err := service.NewSearchProducsService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.HTML(consts.StatusOK, "search", utils.WarpResponse(ctx, c, resp))
}
