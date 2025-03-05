package dal

import (
	"github.com/meilingluolingluo/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
