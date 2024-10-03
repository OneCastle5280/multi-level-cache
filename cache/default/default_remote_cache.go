package _default

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type DefaultRemoteCache struct {
	redisCache redis.Client // default remote cache

}

func (d *DefaultRemoteCache) BatchSet(ctx context.Context, values map[string][]byte, expire time.Duration) error {
	if len(values) == 0 {
		return nil
	}

	// pipeline
	pipeline := d.redisCache.Pipeline()
	for key, val := range values {
		pipeline.Set(ctx, key, val, expire)
	}
	//exec
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultRemoteCache) BatchGet(ctx context.Context, keys []string) (map[string][]byte, []string, error) {
	var notFoundKeys []string
	result := make(map[string][]byte, len(keys))

	if len(keys) == 0 {
		return result, notFoundKeys, nil
	}

	cacheValues, err := d.redisCache.MGet(ctx, keys...).Result()
	if err != nil {
		return result, notFoundKeys, err
	}

	for i := range cacheValues {
		key := keys[i]
		val := cacheValues[i]
		if val == nil {
			notFoundKeys = append(notFoundKeys, key)
		} else {
			if value, ok := val.([]byte); ok {
				result[key] = value
			}
		}
	}

	return result, notFoundKeys, nil
}

func (d *DefaultRemoteCache) BatchDel(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	_, err := d.redisCache.Del(ctx, keys...).Result()
	return err
}
