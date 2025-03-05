package rpc

import (
	"sync"

	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/meilingluolingluo/gomall/app/frontend/conf"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
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
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendutils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulRegister(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts)
	frontendutils.MustHandleError(err)
}
