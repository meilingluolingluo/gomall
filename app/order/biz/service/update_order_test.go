package service

import (
	"context"
	"testing"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func TestUpdateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderService(ctx)
	// init req and assert value
	mysql.InitTest()

	req := &order.UpdateOrderReq{
		OrderId: "0c61bb30-fa90-11ef-a1c0-0242ac110002",
		Address: &order.Address{
			City:          "冬木市",
			Country:       "日本",
			State:         "unknown",
			StreetAddress: "teststreetAddress",
			ZipCode:       10086,
		},
		UserId: 123,
		Email:  "卫宫士郎@qq,com",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
