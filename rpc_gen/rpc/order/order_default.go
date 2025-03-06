package order

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = defaultClient.MarkOrderPaid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MarkOrderPaid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (resp *order.UpdateOrderResp, err error) {
	resp, err = defaultClient.UpdateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteOrder(ctx context.Context, req *order.DeleteOrderReq, callOptions ...callopt.Option) (resp *order.DeleteOrderResp, err error) {
	resp, err = defaultClient.DeleteOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
