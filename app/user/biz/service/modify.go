package service

import (
	"context"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
)

type ModifyService struct {
	ctx context.Context
} // NewModifyService new ModifyService
func NewModifyService(ctx context.Context) *ModifyService {
	return &ModifyService{ctx: ctx}
}

// Run create note info
func (s *ModifyService) Run(req *user.ModifyReq) (resp *user.ModifyResp, err error) {
	// Finish your business logic.

	return
}
