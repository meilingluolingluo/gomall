package dal

import (
	"github.com/meilingluolingluo/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	err := mysql.Init(mysql.DefaultConfig())
	if err != nil {
		return
	}
}
