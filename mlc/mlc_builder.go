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
		localCache:  cache.NewLocalCache[T](config),
		remoteCache: cache.NewRemoteCache[T](config),
		getFromDb:   getFromDb,
		unionKey:    name,
	}
}
