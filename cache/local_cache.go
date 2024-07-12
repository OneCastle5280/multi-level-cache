package cache

import "context"

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
	return &LocalCache[T]{
		cache:       config.getLocalCache(),
		CommonCache: NewCommonCache[T](loader, config),
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

	result := make(map[string][]byte, len(keys))

	if len(keys) == 0 {
		return result, nil
	}

	// batch query from cache
	cacheValueMap, notFoundKeys, err := l.cache.BatchGet(ctx, keys)
	if err != nil {
		// todo 打印日志，查询出现了异常，需要降级回源
		sourceValueMap, loaderErr := l.loader(ctx, keys)
		if loaderErr != nil {
			// todo 打印日志，回源也出现了异常
			return nil, err
		}

		// return source value
		return sourceValueMap, nil
	}

	// add cache value to result
	for key, val := range cacheValueMap {
		result[key] = val
	}

	// handle not found keys
	if len(notFoundKeys) > 0 {
		// loader source value
		sourceValueMap, loaderErr := l.loader(ctx, notFoundKeys)
		if loaderErr != nil {
			// todo 打印日志，回源也出现了异常
			return nil, err
		}

		// add source value to result
		for key, val := range sourceValueMap {
			result[key] = val
		}
	}

	// todo 缓存穿透情况处理

	return result, nil
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
