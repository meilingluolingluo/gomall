package dal

import (
	"github.com/meilingluolingluo/gomall/app/product/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
