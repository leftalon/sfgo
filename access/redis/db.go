package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func GetDB() (client redis.UniversalClient, err error) {
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
	statusCmd := client.Ping(context.TODO())
	return client, statusCmd.Err()
}
