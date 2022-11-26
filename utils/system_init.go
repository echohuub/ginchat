package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	RED *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config app:", viper.Get("db"))
}

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(sqlite.Open(viper.GetString("db.dsn")), &gorm.Config{Logger: newLogger})
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println(user)
}

func InitRedis() {
	RED = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.pass"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	//pong, err := RED.Ping().Result()
	//if err != nil {
	//	fmt.Println("redis init fail", err)
	//	panic(err)
	//} else {
	//	fmt.Println("redis init success", pong)
	//}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish....", msg)
	err = RED.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅消息到Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RED.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe....", msg.Payload)
	return msg.Payload, err
}
