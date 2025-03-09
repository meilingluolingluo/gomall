package service

import (
	"context"
	"testing"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	mysql.InitTest()
	// init req and assert value
	req := &order.PlaceOrderReq{
		UserId:       123,
		UserCurrency: "RMB",
		Email:        "金闪闪@example.com",
		OrderItems: []*order.OrderItem{
			{
				Item: &cart.CartItem{
					ProductId: 10086,
					Quantity:  2001,
				},
				Cost: 19.99,
			},
			{
				Item: &cart.CartItem{
					ProductId: 10057,
					Quantity:  20060,
				},
				Cost: 189.99,
			},
		},
		Address: &order.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "USA",
			ZipCode:       12345,
		},
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
