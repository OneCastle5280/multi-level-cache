package cache

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
func NewLocalCache[T any](config *Config) *LocalCache[T] {
	if config.GetMode() == REMOTE {
		return nil
	}

	return &LocalCache[T]{
		cache:       config.GetLocalCache(),
		CommonCache: NewCommonCache[T](config),
	}
}
