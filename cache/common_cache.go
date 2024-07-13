package cache

import (
	"context"
	"mlc/util"
	"time"
)

// CommonCache
// @Description: 本地、远端缓存公共处理逻辑
type CommonCache[T any] struct {
	//
	//  expire
	//  @Description: 缓存过期时间
	//
	expire time.Duration

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

func NewCommonCache[T any](loader Loader, expire int, config *Config) CommonCache[T] {
	return CommonCache[T]{
		expire:       time.Duration(expire),
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

	// Mark whether there is an exception in this query
	existErr := false
	var breakDownKeys []string

	// batch query from cache
	cacheValueMap, notFoundKeys, err := cache.BatchGet(ctx, keys)
	if err != nil {
		existErr = true
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
		// todo 判断是否为击穿节点

		result[key] = val
	}

	// handle not found keys
	if len(notFoundKeys) > 0 {
		// loader source value
		sourceValueMap, reloadErr := c.reload(ctx, cache, notFoundKeys)
		if reloadErr != nil {
			// todo 打印日志，回源也出现了异常
			return nil, err
		}

		// add source value to result
		for key, val := range sourceValueMap {
			result[key] = val
		}
	}

	// todo 缓存穿透情况处理
	if !existErr {

	}

	return result, nil
}

// reload
//
//	@Description: get from source and reload cache
//	@receiver c
//	@param ctx
//	@param cache
//	@param keys
//	@return map[string][]byte
//	@return error
func (c *CommonCache[T]) reload(ctx context.Context, cache Cache, keys []string) (map[string][]byte, error) {

	if len(keys) == 0 {
		return make(map[string][]byte, len(keys)), nil
	}

	// load from source
	sourceValues, err := c.loader(ctx, keys)
	if err != nil {
		// todo 打印日志，回源出现了异常
		return nil, err
	}

	// reset to cache
	err = cache.BatchSet(ctx, sourceValues, c.expire)
	if err != nil {
		// todo 打印日志，说明 reload 缓存出现了异常
	}

	return sourceValues, nil
}

// collectBreakDownKeys
//
//	@Description: collect break down keys
//	@receiver c
//	@param ctx
//	@param queryKeys
//	@param resultKeys
//	@param breakDownKeys
//	@return []string
func (c *CommonCache[T]) collectBreakDownKeys(ctx context.Context, queryKeys []string, resultKeys []string, breakDownKeys []string) []string {
	if len(queryKeys) == len(resultKeys)+len(breakDownKeys) {
		return []string{}
	}

	var existKeys []string
	existKeys = append(existKeys, resultKeys...)
	existKeys = append(existKeys, breakDownKeys...)

	var result []string
	for _, key := range queryKeys {
		if !util.Contains(existKeys, key) {
			result = append(result, key)
		}
	}
	return result
}
