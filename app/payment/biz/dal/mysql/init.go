package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/meilingluolingluo/gomall/app/cart/biz/model"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global GORM database instance.
var DB *gorm.DB

// Config holds MySQL configuration.
type Config struct {
	DSN               string
	MaxIdleConns      int
	MaxOpenConns      int
	ConnMaxLifetime   time.Duration
	SlowThreshold     time.Duration
	LogLevel          logger.LogLevel
	Colorful          bool
	IgnoreNotFound    bool
	PrepareStmt       bool
	SkipTransaction   bool
	AutoMigrateModels []interface{}
}

// DefaultConfig provides default MySQL configuration for development.
func DefaultConfig() *Config {
	return &Config{
		MaxIdleConns:    10,  // 最大空闲连接数
		MaxOpenConns:    100, // 最大打开连接数
		ConnMaxLifetime: time.Hour,
		SlowThreshold:   200 * time.Millisecond, // 慢查询阈值
		LogLevel:        logger.Info,            // 开发环境显示所有 SQL
		Colorful:        true,                   // 启用彩色日志
		IgnoreNotFound:  true,                   // 忽略 "record not found" 错误
		PrepareStmt:     true,                   // 预编译语句
		SkipTransaction: true,                   // 跳过默认事务
		AutoMigrateModels: []interface{}{
			&model.Cart{},
			// 添加其他需要自动迁移的模型
		},
	}
}

// getDSN generates the MySQL DSN from environment variables.
func getDSN() string {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// 验证所有必需的环境变量
	if mysqlUser == "" {
		klog.Fatalf("MYSQL_USER environment variable is required")
	}
	if mysqlPassword == "" {
		klog.Fatalf("MYSQL_PASSWORD environment variable is required")
	}
	if mysqlHost == "" {
		klog.Fatalf("MYSQL_HOST environment variable is required")
	}
	if mysqlDatabase == "" {
		klog.Fatalf("MYSQL_DATABASE environment variable is required")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
}

// Init initializes the MySQL connection with GORM.
func Init(cfg *Config) error {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	// 获取 DSN
	dsn := getDSN()
	// 配置 GORM 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 打开数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		PrepareStmt:            cfg.PrepareStmt,
		SkipDefaultTransaction: cfg.SkipTransaction,
	})
	if err != nil {
		klog.Errorf("Failed to connect to MySQL: %v", err)
		return fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// 配置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		klog.Errorf("Failed to get SQL DB: %v", err)
		return fmt.Errorf("failed to get SQL DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// 自动迁移表结构
	if err := DB.AutoMigrate(cfg.AutoMigrateModels...); err != nil {
		klog.Errorf("Failed to auto-migrate models: %v", err)
		return fmt.Errorf("failed to auto-migrate models: %v", err)
	}

	klog.Infof("MySQL connection initialized successfully with DSN: %s", dsn)
	return nil
}

// customLogger integrates GORM logs with Kitex klog.
type customLogger struct {
	w logger.Writer
}

func (l *customLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	// 将 ctx 和日志信息合并为一个字符串
	message := fmt.Sprintf(format, v...)
	if ctx != nil {
		klog.Infof("Context: %v, Message: %s", ctx, message)
	} else {
		klog.Infof("Message: %s", message)
	}
	l.w.Printf(format, v...)
}
