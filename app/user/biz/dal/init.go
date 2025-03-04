package dal

import (
	"fmt"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/user/biz/model"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

// 提取 DSN 生成逻辑到独立函数
func getDSN() string {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if mysqlUser == "" || mysqlPassword == "" || mysqlHost == "" {
		panic("missing required environment variables: MYSQL_USER, MYSQL_PASSWORD, or MYSQL_HOST")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
}

// Init 开启数据库连接
func Init() {
	dsn := getDSN()
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	//自动迁移数据库
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		panic(fmt.Errorf("failed to auto migrate database: %w", err))
	}
}
