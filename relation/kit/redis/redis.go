package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func Redis_init() *redis.Client {
	cli := redis.NewClient(
		&redis.Options{
			Addr:         viper.GetString("redis.addr"),
			Password:     viper.GetString("redis.password"),
			DB:           viper.GetInt("redis.DB"),
			PoolSize:     viper.GetInt("redis.PoolSize"),
			MinIdleConns: viper.GetInt("redis.MinIdleConn"),
		},
	)
	Redis = cli
	return Redis
}
