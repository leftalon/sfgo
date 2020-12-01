package redis

import (
	"github.com/go-redis/redis/v8"
)

func GetRedisClient() (client redis.UniversalClient) {
	switch RedisConfig.Mode {
	case "single":
		client = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     RedisConfig.Addr,
			Password: RedisConfig.Password,
			DB:       RedisConfig.DbName,
		})
	case "cluster":
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    RedisConfig.Addrs,
			Password: RedisConfig.Password,
		})
	case "sentinel":
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    RedisConfig.Master,
			SentinelAddrs: RedisConfig.Addrs,
			Password:      RedisConfig.Password,
			DB:            RedisConfig.DbName,
		})
	default:
		client = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     RedisConfig.Addr,
			Password: RedisConfig.Password,
			DB:       RedisConfig.DbName,
		})
	}

	return client
}
