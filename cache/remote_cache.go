package cache

type (
	RemoteCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  Cache
		//  @Description: 远程缓存 Cache
		//
		Cache Cache
	}
)

// NewRemoteCache
//
//	@Description: 创建远程缓存
//	@param mode
//	@return Cache.RemoteCache[T]
func NewRemoteCache[T any](loader Loader, config *Config) *RemoteCache[T] {
	return &RemoteCache[T]{
		Cache:       config.GetRemoteCache(),
		CommonCache: NewCommonCache[T](loader, config),
	}
}
