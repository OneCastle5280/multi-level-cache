package _default

import (
	"context"
	"errors"
	"github.com/coocood/freecache"
	"time"
)

type DefaultLocalCache struct {
	localCache *freecache.Cache // default localCache
}

// NewDefaultLocalCache
//
//	@Description: 创建 default local cache
//	@param localCacheSize
//	@return *DefaultLocalCache
func NewDefaultLocalCache(localCacheSize int32) *DefaultLocalCache {
	return &DefaultLocalCache{
		localCache: freecache.NewCache(int(localCacheSize))}
}

func (d *DefaultLocalCache) BatchSet(ctx context.Context, values map[string][]byte, expire time.Duration) error {
	if len(values) == 0 {
		return nil
	}

	for key, val := range values {
		err := d.localCache.Set([]byte(key), val, int(expire.Seconds()))
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DefaultLocalCache) BatchGet(ctx context.Context, keys []string) (map[string][]byte, []string, error) {
	var notFoundKeys []string
	result := make(map[string][]byte, len(keys))

	if len(keys) == 0 {
		return result, notFoundKeys, nil
	}

	for _, key := range keys {
		val, err := d.localCache.Get([]byte(key))
		if err != nil {
			if errors.Is(err, freecache.ErrNotFound) {
				notFoundKeys = append(notFoundKeys, key)
			} else {
				return result, notFoundKeys, err
			}
		} else {
			result[key] = val
		}
	}

	return result, notFoundKeys, nil

}

func (d *DefaultLocalCache) BatchDel(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	for _, key := range keys {
		d.localCache.Del([]byte(key))
	}

	return nil
}
