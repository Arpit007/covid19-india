package cache

import (
	"context"
	"covid19-india/internal/config"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

var redisClient *redis.Client

// RedisCache Config of a cache instance
type RedisCache struct {
	ttl    time.Duration
	prefix string
}

type FetchCallback func(ctx context.Context) (interface{}, error)

func init() {
	u, err := url.Parse(config.ENV.RedisUri)

	if err != nil {
		logrus.Fatal("Error parsing redis uri ", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Username: u.User.Username(),
		Password: getPasswordFromURL(u),
		DB:       0,
	})

	if _, err := redisClient.Ping(context.TODO()).Result(); err != nil {
		logrus.Fatal("Unable to connect to Redis: ", err)
	}
}

func getPasswordFromURL(url *url.URL) string {
	p, isSet := url.User.Password()

	if isSet {
		return p
	}

	return ""
}

// CreateRedisCache Create a new redis cache
func CreateRedisCache(ttl time.Duration, prefix string) *RedisCache {
	return &RedisCache{
		ttl:    ttl,
		prefix: prefix,
	}
}

// Get Gets value from cache, else uses callback in case of cache miss & updates the cache
func (rCache *RedisCache) Get(ctx context.Context, key string, v interface{}, callback FetchCallback) (interface{}, error) {
	rKey := rCache.prepareKey(key)

	// Fetch from redis
	if res, err := redisClient.Get(ctx, rKey).Bytes(); err == nil {
		if err := json.Unmarshal(res, v); err == nil {
			return v, nil
		}
	}

	data, err := callback(ctx) // Cache miss, use callback to get value

	if err != nil {
		return nil, err
	}

	go rCache.Set(context.TODO(), key, data) // update the cache

	return data, nil
}

// Prepare redis cache key
func (rCache *RedisCache) prepareKey(key string) string {
	return rCache.prefix + ":" + key // To avoid collision in keys of different caches
}

// Set value in the cache
func (rCache *RedisCache) Set(ctx context.Context, key string, v interface{}) {
	rKey := rCache.prepareKey(key)

	data, err := json.Marshal(v)

	if err != nil {
		logrus.Error("Failed to marshal data")
		return
	}

	if _, err := redisClient.SetEX(ctx, rKey, data, rCache.ttl).Result(); err != nil {
		logrus.Error("Failed to update cache item ", key)
	}
}

// RemoveKeys Remove keys from cache
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
