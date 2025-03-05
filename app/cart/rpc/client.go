package rpc

import (
	cartutils "github.com/meilingluolingluo/gomall/app/cart/utils"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/meilingluolingluo/gomall/app/cart/conf"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	once          sync.Once
	err           error
	commonSuite   client.Option
	ProductClient productcatalogservice.Client
)

func Init() {
	once.Do(func() {
		//initUserClient()
		initProductClient()
		//initCartClient()
	})
}

func initClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	cartutils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("cart", opts...)
	cartutils.MustHandleError(err)
}
