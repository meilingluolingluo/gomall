package category

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/service"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/utils"
	category "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/category"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "category", resp)
}
