package dal

import (
	"github.com/meilingluolingluo/gomall/app/frontend/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
