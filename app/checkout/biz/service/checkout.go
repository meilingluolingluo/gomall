package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/meilingluolingluo/gomall/app/checkout/infra/rpc"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	//获取用户购物车信息
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	// 遍历购物车信息
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)
	// charge
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	// 根据支付类型处理支付请求
	var payReq *payment.ChargeReq
	switch req.Payment.PaymentType { // 根据checkout.proto使用PaymentType枚举
	case checkout.PaymentType_ALIPAY:
		payReq = &payment.ChargeReq{
			UserId:        req.UserId,
			OrderId:       orderId,
			Amount:        total,
			PaymentMethod: payment.PaymentMethod_ALIPAY, // 转换为payment服务的枚举
			PaymentInfo: &payment.ChargeReq_AlipayAuthCode{
				AlipayAuthCode: req.Payment.AlipayAuthCode, // 直接访问字段
			},
		}
	default: // 默认信用卡支付
		payReq = &payment.ChargeReq{
			UserId:        req.UserId,
			OrderId:       orderId,
			Amount:        total,
			PaymentMethod: payment.PaymentMethod_CREDIT_CARD,
			PaymentInfo: &payment.ChargeReq_CreditCard{
				CreditCard: &payment.CreditCardInfo{
					CreditCardNumber:          req.Payment.GetCreditCard().CreditCardNumber, // 使用Get方法访问嵌套字段
					CreditCardCvv:             req.Payment.GetCreditCard().CreditCardCvv,
					CreditCardExpirationYear:  req.Payment.GetCreditCard().CreditCardExpirationYear,
					CreditCardExpirationMonth: req.Payment.GetCreditCard().CreditCardExpirationMonth,
				},
			},
		}
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
		PaymentUrl:    paymentResult.PaymentUrl,
	}
	return
}
