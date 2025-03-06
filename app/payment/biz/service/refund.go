package service

import (
	"context"
	payment "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/payment"
)

type RefundService struct {
	ctx context.Context
} // NewRefundService new RefundService
func NewRefundService(ctx context.Context) *RefundService {
	return &RefundService{ctx: ctx}
}

// Run create note info
func (s *RefundService) Run(req *payment.RefundReq) (resp *payment.RefundResp, err error) {
	// Finish your business logic.

	return
}
