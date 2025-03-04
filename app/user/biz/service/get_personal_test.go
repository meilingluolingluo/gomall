package service

import (
	"context"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestGetPersonal_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetPersonalService(ctx)
	// init req and assert value

	req := &user.GetPersonalReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
