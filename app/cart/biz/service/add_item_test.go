package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/meilingluolingluo/gomall/app/cart/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/cart/infra/rpc"
	cart "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Logf("err: %v", err)
	}
	err = mysql.Init(mysql.DefaultConfig())
	if err != nil {
		return
	}

	ctx := context.Background()
	s := NewGetCartService(ctx)

	// init req and assert value
	rpc.Init()

	req := &cart.GetCartReq{
		UserId: 1,
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
