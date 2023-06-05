package redis

import (
	"blog_server/utils/config"
	"time"

	"github.com/go-redis/redis"
)

var instance *redis.Client

func init() {
	instance = redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
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
