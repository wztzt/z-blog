package utils

import (
	"blog_server/utils/config"
	"blog_server/utils/redis"
)

func Start() {
	config.Start()
	redis.Start()
}

func Close() {
	redis.Close()

	config.Close()
}
