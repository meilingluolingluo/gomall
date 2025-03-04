package mysql

import (
	"fmt"
	"github.com/meilingluolingluo/gomall/app/user/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func Init() {
	// 测试环境时
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	dsn := getDSN()
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	err := DB.AutoMigrate( //nolint:errcheck
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
	//DB.Exec("INSERT INTO `user` (`id`,`created_at`,`updated_at`,`username`,`email`,`password_hashed`) VALUES (1,'2025-02-16 09:46:19.852','2025-02-16 09:46:19.852','管理员','123@admin.com','$2a$10$SYSi3FMYYDHDkty.yeqHz.Ft/aoAbOLS9O72MEIsdy/wPEsYl8P8e')")

}
