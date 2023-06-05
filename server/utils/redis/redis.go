package redis

import (
	"blog_server/utils/config"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var instance *redis.Client

func init() {
	aa, _ := redis.ParseURL("redis://red-cg7cfn5269v5l610obtg:6379")
	fmt.Printf("%+v\n", aa)
	options := &redis.Options{
		Addr:     config.RedisAddr,
		PoolSize: config.RedisPool,
	}
	fmt.Printf("%+v\n", options)
	instance = redis.NewClient(options)

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
