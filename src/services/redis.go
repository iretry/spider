package services

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}
	addr := viper.GetString("REDIS_ADDR")
	password := viper.GetString("REDIS_PASSWORD")
	db := viper.GetInt("REDIS_DB")
	RedisClient := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password, // no password set
		DB:          db,       // use default DB
		DialTimeout: 1 * time.Second,
	})
	return RedisClient
}
