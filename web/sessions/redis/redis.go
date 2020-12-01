package redis

import (
	"context"
	"errors"
	"sfgo/web/sessions"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

type Store interface {
	sessions.Store
}

func NewStore(client redis.UniversalClient) (Store, error) {
	s, err := redisstore.NewRedisStore(context.TODO(), client)
	if err != nil {
		return nil, err
	}
	s.KeyPrefix("session_")
	return &store{s}, nil
}

type store struct {
	*redisstore.RedisStore
}

func GetRedisStore(s Store) (err error, rediStore *redisstore.RedisStore) {
	realStore, ok := s.(*store)
	if !ok {
		err = errors.New("unable to get the redis store: Store isn't *store")
		return
	}

	rediStore = realStore.RedisStore
	return
}

// SetKeyPrefix sets the key prefix in the redis database.
func SetKeyPrefix(s Store, prefix string) error {
	err, rediStore := GetRedisStore(s)
	if err != nil {
		return err
	}

	rediStore.KeyPrefix(prefix)
	return nil
}

func (c *store) Options(options sessions.Options) {
	c.RedisStore.Options(*options.ToGorillaOptions())
}
