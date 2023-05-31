package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Domain string
	Port   int
)

func init() {
	fmt.Println("init")
	config := viper.New()
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
	Domain = config.GetString("server.domain")
	Port = config.GetInt("server.port")
	env_domain := os.Getenv("domain")
	if env_domain != "" {
		Domain = env_domain
	}
	env_port := os.Getenv("port")
	if env_port != "" {
		Port, _ = strconv.Atoi(env_port)
	}
}
