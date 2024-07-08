package cache

type (
	RemoteCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  cache
		//  @Description: 远程缓存 cache
		//
		cache Cache
	}
)

// NewRemoteCache
//
//	@Description: 创建远程缓存
//	@param mode
//	@return cache.RemoteCache[T]
func NewRemoteCache[T any](config *Config) *RemoteCache[T] {
	if config.GetMode() == LOCAL {
		// 本地缓存模式，无需创建远程缓存
		return nil
	}

	return &RemoteCache[T]{
		cache:       config.GetRemoteCache(),
		CommonCache: NewCommonCache[T](config),
	}
}
