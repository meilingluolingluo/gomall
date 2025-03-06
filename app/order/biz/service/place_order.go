package service

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/order/model"
	order "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		err = kerrors.NewGRPCBizStatusError(2004001, "order items is required")
	}
	log.Printf("using mysql %v", mysql.DB)
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := model.Order{
			OrderId:      orderId.String(),
			UserId:       uint(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}

		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.Country = a.Country
			o.Consignee.State = a.State
		}

		if err := tx.Create(&o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				Quantity:     uint(v.Item.Quantity),
				ProductId:    uint(v.Item.ProductId),
				Cost:         float64(v.Cost),
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	return
}
