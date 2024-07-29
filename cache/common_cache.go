package cache

import (
	"context"
	"mlc/cache/log"
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

	//
	//  breakDownHandler
	//  @Description: 缓存击穿处理器
	//
	breakDownHandler CacheBreakDownHandler
}

// NewCommonCache
//
//	@Description: new common cache
//	@param loader
//	@param expire
//	@param config
//	@return CommonCache[T]
func NewCommonCache[T any](loader Loader, expire int, config *Config) CommonCache[T] {
	breakDownHandler := config.breakDownHandler
	if breakDownHandler == nil {
		breakDownHandler = NewDefaultCacheBreakDownHandler()
	}

	return CommonCache[T]{
		expire:           time.Duration(expire),
		loader:           loader,
		breakDownHandler: breakDownHandler,
		statsHandler:     NewStatsHandler(config.statsDisable, config.statsHandler),
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
		log.Error("[batchGet] query keys: %+v err:%+v, need reload source", keys, err)
		sourceValueMap, loaderErr := c.loader(ctx, keys)
		if loaderErr != nil {
			log.Error("[batchGet] reload keys: %+v err:%+v", keys, err)
			return nil, err
		}

		// return source value
		log.Info("[batchGet] keys: %+v, sourceValueMap: %+v", keys, sourceValueMap)
		return sourceValueMap, nil
	}

	// add cache value to result
	for key, val := range cacheValueMap {
		if c.breakDownHandler.IsBreakDownKeys(nil, val) {
			// exist breakDownKeys
			log.Info("[batchGet] %s is break down key", key)
			breakDownKeys = append(breakDownKeys, key)
		} else {
			result[key] = val
		}
	}

	// handle not found keys
	if len(notFoundKeys) > 0 {
		// loader source value
		sourceValueMap, reloadErr := c.reload(ctx, cache, notFoundKeys)
		if reloadErr != nil {
			log.Error("[batchGet] reload keys: %+v err:%+v", keys, err)
			return nil, err
		}

		// add source value to result
		for key, val := range sourceValueMap {
			result[key] = val
		}
	}

	if !existErr {
		// 查询没有出现异常，处理 缓存穿透场景
		c.handleBreakDownKeys(ctx, cache, keys, util.Keys(result), breakDownKeys)
	}

	log.Info("[batchGet] keys: %+v, result: %+v", keys, result)
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
		return nil, err
	}

	// reload to cache
	err = cache.BatchSet(ctx, sourceValues, c.expire)
	if err != nil {
		// reload 异常则降级，等下次重新写入
		log.Error("[reload] reload sourceValues: %+v to cache err %+v", sourceValues, err)
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
//	@param existBreakDownKeys
//	@return []string
func (c *CommonCache[T]) collectBreakDownKeys(ctx context.Context, queryKeys []string, resultKeys []string, existBreakDownKeys []string) []string {
	if len(queryKeys) == len(resultKeys)+len(existBreakDownKeys) {
		return []string{}
	}

	var existKeys []string
	existKeys = append(existKeys, resultKeys...)
	existKeys = append(existKeys, existBreakDownKeys...)

	var breakDownKeys []string
	for _, key := range queryKeys {
		if !util.Contains(existKeys, key) {
			breakDownKeys = append(breakDownKeys, key)
		}
	}
	log.Info("[collectBreakDownKeys] break down keys: %+v", breakDownKeys)
	return breakDownKeys
}

// handleBreakDownKeys
//
//	@Description: 处理击穿节点
//	@receiver c
//	@param ctx
//	@param queryKeys
//	@param resultKeys
//	@param existBreakDownKeys
func (c *CommonCache[T]) handleBreakDownKeys(ctx context.Context, cache Cache, queryKeys []string, resultKeys []string, existBreakDownKeys []string) {
	// collect need handle break down keys
	needHandleBreakDownKeys := c.collectBreakDownKeys(ctx, queryKeys, resultKeys, existBreakDownKeys)
	if len(needHandleBreakDownKeys) == 0 {
		return
	}
	values := c.breakDownHandler.HandleBreakDownKeys(ctx, needHandleBreakDownKeys)
	log.Info("[handleBreakDownKeys] values: %+v", values)
	err := cache.BatchSet(ctx, values, c.expire)
	if err != nil {
		log.Error("[handleBreakDownKeys] handle break down keys: %+v, err %+v", needHandleBreakDownKeys, err)
	}
}
