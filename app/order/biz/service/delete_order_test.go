package service

import (
	"context"
	"testing"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func TestDeleteOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteOrderService(ctx)
	mysql.InitTest()
	// init req and assert value
	req := &order.DeleteOrderReq{
		UserId:  123,
		OrderId: "a94972d8-fa8f-11ef-addf-0242ac110002",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
