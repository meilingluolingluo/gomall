package rpc

import (
	"github.com/cloudwego/kitex/client"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user/userservice"

	"sync"
)

var (
	UserClient  userservice.Client
	once        sync.Once
	err         error
	commonSuite client.Option
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendutils.MustHandleError(err)
}
