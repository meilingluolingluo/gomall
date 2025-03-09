package order

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/service"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/utils"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/order"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// log.Printf("resp order %v", resp)
	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}

// UpdateOrderPage .
// @router /order/update/:order_id [GET]
func UpdateOrderPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUpdateOrderPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "updateorderpage", utils.WarpResponse(ctx, c, resp))
}

// UpdateOrder .
// @router /order/update/:order_id [POST]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderReq
	err = c.BindAndValidate(&req)
	// log.Printf("UpdateOrderReq %v", req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// log.Printf("UpdateOrderReq %s", req.OrderId)
	_, err = service.NewUpdateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 重定向到 /order 页面
	c.Redirect(consts.StatusFound, []byte("/order"))
}

// DeleteOrder .
// @router /order/delete/:order_id [POST]
func DeleteOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.DeleteOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	log.Printf("DeleteOrderid %s", req.OrderId)
	resp, err := service.NewDeleteOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)

	c.HTML(consts.StatusOK, "delete-order", utils.WarpResponse(ctx, c, resp))
}
