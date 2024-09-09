package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) Get(key string) ([]byte, error) {
	data, err := r.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	serializedValue, err := serializeValue(value)
	if err != nil {
		return err
	}
	return r.client.Set(key, serializedValue, expiration).Err()
}

func serializeValue(value interface{}) ([]byte, error) {
	switch v := value.(type) {
	case string:
		return []byte(v), nil
	default:
		return json.Marshal(value)
	}
}
