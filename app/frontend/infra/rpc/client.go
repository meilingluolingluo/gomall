package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/meilingluolingluo/gomall/app/frontend/conf"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/mtl"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"github.com/meilingluolingluo/gomall/common/clientsuite"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: frontendutils.ServiceName,
		})
		initProductClient()
		initUserClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initProductClient() {
	var opts []client.Option

	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	opts = append(opts, commonSuite, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
		if err == nil {
			return resp, err
		}
		if methodName != "ListProducts" {
			return resp, err
		}
		return &product.ListProductsResp{
			Products: []*product.Product{
				{
					Price:       5.6,
					Id:          3,
					Picture:     "/static/image/t-shirt.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
				{
					Price:       6.6,
					Id:          4,
					Picture:     "/static/image/t-shirt-1.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
				{
					Price:       6.6,
					Id:          4,
					Picture:     "/static/image/t-shirt-1.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
				{
					Price:       6.6,
					Id:          4,
					Picture:     "/static/image/t-shirt-1.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
				{
					Price:       6.6,
					Id:          4,
					Picture:     "/static/image/t-shirt-1.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
				{
					Price:       6.6,
					Id:          4,
					Picture:     "/static/image/t-shirt-1.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
			},
		}, nil
	}))))
	opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	frontendutils.MustHandleError(err)
}
