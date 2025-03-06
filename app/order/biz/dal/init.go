package dal

import (
	"github.com/meilingluolingluo/gomall/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init(mysql.DefaultConfig())
}
