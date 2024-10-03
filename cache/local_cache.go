package cache

import (
	"context"
	_default "mlc/cache/default"
)

type (
	LocalCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  cache
		//  @Description: 本地缓存 cache
		//
		cache Cache
	}
)

// NewLocalCache
//
//	@Description: 创建本地缓存
//	@param config
//	@return *cache.LocalCache[T]
func NewLocalCache[T any](loader Loader, config *Config) *LocalCache[T] {
	cache := config.GetLocalCache()
	if cache == nil {
		// default local cache
		cache = _default.NewDefaultLocalCache(int32(config.GetLocalLimitSize()))
	}
	return &LocalCache[T]{
		cache:       cache,
		CommonCache: NewCommonCache[T](loader, config.getLocalExpire(), config),
	}
}

// BatchGet
//
//	@Description: 批量获取本地缓存
//	@receiver l
//	@param ctx
//	@param keys
//	@return map[string][]byte
//	@return error
func (l *LocalCache[T]) BatchGet(ctx context.Context, keys []string) (map[string][]byte, error) {
	return l.batchGet(ctx, l.cache, l.statsHandler, LocalType, keys)
}

// BatchDel
//
//	@Description: 批量清理本地缓存
//	@receiver l
//	@param ctx
//	@param keys
//	@return error
func (l *LocalCache[T]) BatchDel(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	return l.cache.BatchDel(ctx, keys)
}
