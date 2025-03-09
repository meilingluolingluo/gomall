package service

import (
	"context"
	"log"
	"testing"

	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value
	mysql.InitTest()
	req := &order.ListOrderReq{
		UserId: 13,
	}
	resp, err := s.Run(req)

	log.Printf("resp = : %v", resp)
	t.Logf("err: %v", err)
	t.Logf("resp  = : %v", resp)
}
