package cache

import (
	"context"
	"time"

	"github.com/Ahmad940/assetly_server/pkg/config"
	"github.com/redis/go-redis/v9"
)

// redisClient the client for redis
var redisClient = redis.NewClient(&redis.Options{
	Addr:     config.GetEnv().REDIS_ADDRESS,
	Password: config.GetEnv().REDIS_PASSWORD,
	DB:       0, // use default DB
})

// SetRedisValue SetValue sets the key value pair
func SetRedisValue(key string, value string, expiry time.Duration) error {
	err := redisClient.Set(context.Background(), key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRedisValue GetValue the value corresponding to a given key
func GetRedisValue(key string) (string, error) {
	value, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// DelRedisValue del key
func DelRedisValue(key string) error {
	return redisClient.Del(context.Background(), key).Err()
}
