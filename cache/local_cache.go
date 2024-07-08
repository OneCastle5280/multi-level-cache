package cache

type (
	LocalCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  Cache
		//  @Description: 本地缓存 Cache
		//
		Cache Cache
	}
)

// NewLocalCache
//
//	@Description: 创建本地缓存
//	@param config
//	@return *Cache.LocalCache[T]
func NewLocalCache[T any](loader Loader, config *Config) *LocalCache[T] {
	return &LocalCache[T]{
		Cache:       config.GetLocalCache(),
		CommonCache: NewCommonCache[T](loader, config),
	}
}
