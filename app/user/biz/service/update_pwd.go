package service

import (
	"context"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
)

type UpdatePwdService struct {
	ctx context.Context
} // NewUpdatePwdService new UpdatePwdService
func NewUpdatePwdService(ctx context.Context) *UpdatePwdService {
	return &UpdatePwdService{ctx: ctx}
}

// Run create note info
func (s *UpdatePwdService) Run(req *user.UpdatePwdReq) (resp *user.UpdatePwdResp, err error) {
	// Finish your business logic.

	return
}
