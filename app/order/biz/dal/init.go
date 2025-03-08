package dal

import (
	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	err := mysql.Init(mysql.DefaultConfig())
	if err != nil {
		return
	}
}
