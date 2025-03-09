package service

import (
	"context"
	"testing"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func TestMarkOrderPaid_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)
	// init req and assert value
	mysql.InitTest()

	req := &order.MarkOrderPaidReq{
		OrderId: "3618adf7-fb06-11ef-a3f2-0242ac110002",
		UserId:  123,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
