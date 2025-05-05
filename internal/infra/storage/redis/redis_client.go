package redis_cache

import "github.com/redis/go-redis/v9"

// stub for redis cache, implement is enough time
type RedisClient struct {
	rc *redis.Client
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		rc: redis.NewClient(&redis.Options{}),
	}
}

func (rc *RedisClient) Get(string) (string, error) {
	return "", nil
}

func (rc *RedisClient) Set(key string, value string) error {
	return nil
}
