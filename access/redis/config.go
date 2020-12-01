package redis

import "sfgo/core/config"

var RedisConfig struct {
	Addr     string
	Password string
	Addrs    []string
	DbName   int
	Mode     string
	Master   string
}

func init() {
	RedisConfig.Addr = "localhost:6379"
	RedisConfig.Mode = "single" //single, sentinel, cluster
	RedisConfig.Master = "master"

	config.Register("sfgo.db.redis", &RedisConfig)
}
