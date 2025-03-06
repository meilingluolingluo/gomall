package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/meilingluolingluo/gomall/app/checkout/conf"
	checkoututils "github.com/meilingluolingluo/gomall/app/checkout/utils"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
	commonSuite   client.Option
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	checkoututils.MustHandleError(err)
}
