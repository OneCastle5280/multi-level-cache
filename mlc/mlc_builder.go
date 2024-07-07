package mlc

import (
	"mlc/cache"
)

var (
	// 记录所有 mlc 实例
	mlcMap = make(map[string]struct{})
)

// NewDefaultMultiLevelCache
//
//	@Description: 创建一个默认的多级缓存实例
//	@param getFromDb		回源接口
//	@param name				多级缓存实例名称
//	@param opt				可选配置
//	@return *DefaultMultiLevelCache[T]
func NewDefaultMultiLevelCache[T any](getFromDb cache.Loader, name string, opt ...cache.ConfigOption) *DefaultMultiLevelCache[T] {

	// 检查 name 唯一性
	if _, ok := mlcMap[name]; ok {
		panic(name + "had existed, please check!")
	}

	// 缓存配置
	config := cache.NewCacheConfig(opt...)

	// 创建缓存实例
	return &DefaultMultiLevelCache[T]{
		config:      config,
		localCache:  initLocalCache[T](config),
		remoteCache: initRemoteCache[T](config),
		getFromDb:   getFromDb,
		unionKey:    name,
	}
}

// initRemoteCache[T any]
//
//	@Description: 创建远程缓存
//	@param mode
//	@return cache.RemoteCache[T]
func initRemoteCache[T any](config *cache.Config) *cache.RemoteCache[T] {
	if config.GetMode() == cache.LOCAL {
		// 本地缓存模式，无需创建远程缓存
		return nil
	}

	return cache.NewRemoteCache[T](config)
}

// initLocalCache[T any]
//
//	@Description: 创建本地缓存
//	@param config
//	@return *cache.LocalCache[T]
func initLocalCache[T any](config *cache.Config) *cache.LocalCache[T] {
	if config.GetMode() == cache.REMOTE {
		return nil
	}

	return cache.NewLocalCache[T](config)
}
