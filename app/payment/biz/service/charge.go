package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/meilingluolingluo/gomall/app/payment/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/payment/biz/model"
	payment "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/payment"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) handleCreditCard(req *payment.ChargeReq) (*payment.ChargeResp, error) {
	card := creditcard.Card{
		Number: req.GetCreditCard().CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.GetCreditCard().CreditCardCvv)),
		Month:  strconv.Itoa(int(req.GetCreditCard().CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.GetCreditCard().CreditCardExpirationYear)),
	}

	if err := card.Validate(true); err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: translationId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
}
