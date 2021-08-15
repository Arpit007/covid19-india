package cache

import (
	"context"
	"covid19-india/internal/config"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.ENV.RedisUri,
		Password: config.ENV.RedisPassword,
		DB:       0,
	})

	if _, err := redisClient.Ping(context.TODO()).Result(); err != nil {
		logrus.Fatal("Unable to connect to Redis: " + err.Error())
	}
}

type RedisCache struct {
	ttl    time.Duration
	prefix string
}

type FetchCallback func() (interface{}, error)

func CreateRedisCache(ttl time.Duration, prefix string) *RedisCache {
	return &RedisCache{
		ttl:    ttl,
		prefix: prefix,
	}
}

func (rCache *RedisCache) Get(ctx context.Context, key string, v interface{}, callback FetchCallback) (interface{}, error) {
	rKey := rCache.prepareKey(key)

	if res, err := redisClient.Get(ctx, rKey).Bytes(); err == nil {
		if err := json.Unmarshal(res, v); err == nil {
			return v, nil
		}
	}

	data, err := callback()

	if err != nil {
		return nil, err
	}

	rCache.set(ctx, rKey, data)

	return data, nil
}

func (rCache *RedisCache) prepareKey(key string) string {
	return rCache.prefix + ":" + key
}

func (rCache *RedisCache) set(ctx context.Context, key string, v interface{}) {
	data, err := json.Marshal(v)

	if err != nil {
		logrus.Error("Failed to marshal data")
		return
	}

	if _, err := redisClient.SetEX(ctx, key, data, rCache.ttl).Result(); err != nil {
		logrus.Error("Failed to update cache item " + key)
	}
}

func (rCache *RedisCache) RemoveKeys(ctx context.Context, keys []string) error {
	var rKeys []string

	for _, key := range keys {
		rKeys = append(rKeys, rCache.prepareKey(key))
	}

	if _, err := redisClient.Del(ctx, rKeys...).Result(); err != nil {
		return err
	}

	return nil
}
