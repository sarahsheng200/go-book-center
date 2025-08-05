package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-book-center/app/common"
	conf "go-book-center/app/config"
	"time"
)

var logger = common.Logger
var Redis *redis.Client

func init() {
	ctx := context.Background()
	config := conf.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Errorf("redis ping error: %s", err.Error())
		return
	}
	logger.Infof("redis ping: %s", ping)

	Redis = client
}
