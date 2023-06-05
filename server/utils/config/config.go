package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config    *viper.Viper
	Port      int
	RedisAddr string
	RedisPool int
)

func init() {
	fmt.Println("init")
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./conf")
	config.ReadInConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("fs change")
		parse(config)
	})
	config.WatchConfig()
	parse(config)
}

func parse(config *viper.Viper) {
	Port = config.GetInt("server.port")
	RedisAddr = config.GetString("redis.addr")
	RedisPool = config.GetInt("redis.poolsize")
	env_port := os.Getenv("port")
	if env_port != "" {
		Port, _ = strconv.Atoi(env_port)
	}
	redis_addr := os.Getenv("redis_addr")
	if redis_addr != "" {
		RedisAddr = redis_addr
	}
	redis_pool := os.Getenv("redis_pool")
	if redis_pool != "" {
		RedisPool, _ = strconv.Atoi(redis_pool)
	}
}

func GetInt(key string) int {
	return config.GetInt(key)
}

func GetString(key string) string {
	return config.GetString(key)
}
