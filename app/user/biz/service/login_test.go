package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return
	}
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	req := &user.LoginReq{
		Email:    "2538@test.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
