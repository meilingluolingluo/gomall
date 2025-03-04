package dal

import (
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
