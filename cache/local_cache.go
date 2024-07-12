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

	cacheValueMap, err := l.cache.BatchGet(ctx, keys)

	if err != nil {
		// todo 日志打印
		return nil, err
	}

	return cacheValueMap, nil
}
