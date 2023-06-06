package redis

import (
	"blog_server/utils/config"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var instance *redis.Client

func init() {
	options := &redis.Options{
		Addr:         config.RedisAddr,
		PoolSize:     config.RedisPool,
		MinIdleConns: config.RedisIdle,
	}
	instance = redis.NewClient(options)
	result, err := instance.Ping().Result()
	if err == nil {
		fmt.Printf("init redis succuss, result = %v, options = %+v\n", result, options)
	} else {
		fmt.Printf("init redis faild, error = %v\n", err.Error())
	}

}

func Info() *redis.StringCmd {
	return instance.Info()
}

func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return instance.Set(key, value, expiration)
}

func Get(key string) *redis.StringCmd {
	return instance.Get(key)
}

func Exists(keys ...string) *redis.IntCmd {
	return instance.Exists(keys...)
}
