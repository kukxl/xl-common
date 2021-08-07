package container

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"xl-common/pkg/consul"
	"xl-common/pkg/logger"
)

var (
	DB 			*gorm.DB
	RedisCli 	*redis.Client
)

func InitRedis() {
	conf := consul.KVContext
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB: conf.Redis.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Logger.Panic("ping redis fail", zap.Error(err))
	}
	if pong != "PONG" {
		logger.Logger.Panic("客户端连接redis服务端失败")
	}
	RedisCli = client
	logger.Logger.Info("init redis success")
}

func InitMysql() {
	dateSource := consul.KVContext.DateSource
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dateSource.UserName, dateSource.Password, dateSource.Host, dateSource.Port, dateSource.Schema)
	newLogger := gormLog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLog.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      gormLog.Error, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Logger.Panic("客户端连接redis服务端失败")
	}
	DB = db
	logger.Logger.Info("init mysql success")
}
