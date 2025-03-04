package service

import (
	"context"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
)

type GetPersonalService struct {
	ctx context.Context
} // NewGetPersonalService new GetPersonalService
func NewGetPersonalService(ctx context.Context) *GetPersonalService {
	return &GetPersonalService{ctx: ctx}
}

// Run create note info
func (s *GetPersonalService) Run(req *user.GetPersonalReq) (resp *user.GetPersonalResp, err error) {
	// Finish your business logic.

	return
}
