package dal

import (
	"github.com/meilingluolingluo/gomall/app/checkout/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
