package cache

import "context"

// CommonCache
// @Description: 本地、远端缓存公共处理逻辑
type CommonCache[T any] struct {
	//
	//  loader
	//  @Description: 回源接口
	//
	loader Loader

	//
	//  statsHandler
	//  @Description: 统计组件
	//
	statsHandler *StatsHandler
}

func NewCommonCache[T any](loader Loader, config *Config) CommonCache[T] {
	return CommonCache[T]{
		loader:       loader,
		statsHandler: NewStatsHandler(config.statsDisable, config.statsHandler),
	}
}

// batchGet
//
//	@Description: 公共逻辑：批量获取缓存数据
//	@receiver c
//	@param ctx
//	@param cache
//	@param keys
//	@return map[string][]byte
//	@return error
func (c *CommonCache[T]) batchGet(ctx context.Context, cache Cache, keys []string) (map[string][]byte, error) {
	result := make(map[string][]byte, len(keys))

	if len(keys) == 0 {
		return result, nil
	}

	// batch query from cache
	cacheValueMap, notFoundKeys, err := cache.BatchGet(ctx, keys)
	if err != nil {
		// todo 打印日志，查询出现了异常，需要降级回源
		sourceValueMap, loaderErr := c.loader(ctx, keys)
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
		sourceValueMap, loaderErr := c.loader(ctx, notFoundKeys)
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
