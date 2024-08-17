package mlc

import (
	"context"
	"mlc/cache"
	"sync"
)

var (
	// 记录所有 mlc 实例
	mlcMap      = make(map[string]struct{})
	builderLock sync.Mutex
)

// NewDefaultMultiLevelCache
//
//	@Description: 创建一个默认的多级缓存实例
//	@param getFromDb		回源接口
//	@param name				多级缓存实例名称
//	@param opt				可选配置
//	@return *DefaultMultiLevelCache[T]
func NewDefaultMultiLevelCache[T any](getFromDb cache.Loader, name string, opt ...cache.ConfigOption) *DefaultMultiLevelCache[T] {
	builderLock.Lock()
	defer builderLock.Unlock()

	// 检查 name 唯一性
	if _, ok := mlcMap[name]; ok {
		panic(name + "had existed, please check!")
	}

	// 缓存配置
	config := cache.NewCacheConfig(opt...)

	var localCache *cache.LocalCache[T]
	var remoteCache *cache.RemoteCache[T]
	var localCacheLoader cache.Loader

	// 根据缓存模式进行初始化
	switch config.GetMode() {
	case cache.LOCAL:
		localCache = cache.NewLocalCache[T](getFromDb, config)
		break
	case cache.REMOTE:
		remoteCache = cache.NewRemoteCache[T](getFromDb, config)
		break
	case cache.MULTILEVEL:
		remoteCache = cache.NewRemoteCache[T](getFromDb, config)
		localCacheLoader = func(ctx context.Context, keys []string) (map[string][]byte, error) {
			return remoteCache.BatchGet(ctx, keys)
		}
		localCache = cache.NewLocalCache[T](localCacheLoader, config)
	}

	serialization := config.GetSerialization()
	if serialization == nil {
		// default json serialization
		serialization = cache.NewJsonSerialization()
	}

	statsHandler := cache.NewStatsHandler(config.GetStatsDisable(), config.GetStatsHandler())

	// 创建缓存实例
	return &DefaultMultiLevelCache[T]{
		config:        config,
		localCache:    localCache,
		remoteCache:   remoteCache,
		getFromDb:     getFromDb,
		serialization: serialization,
		statsHandler:  statsHandler,
		unionKey:      name,
	}
}
