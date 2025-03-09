package service

import (
	"context"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"testing"
)

func TestDeleteOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteOrderService(ctx)
	// init req and assert value

	req := &order.DeleteOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
